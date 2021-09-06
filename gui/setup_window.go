package gui

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"mnimidamonbackend/adapter/restapi"
	"os"
	"path/filepath"
)

type SetupWindow struct {
	Content fyne.CanvasObject

	Port        binding.Int
	FolderEntry binding.String
}

func (s *SetupWindow) GetContent() fyne.CanvasObject {
	return s.Content
}

func NewSetupWindow(gi *GraphicalInterface, nextWindow ContentGetter) *SetupWindow {
	pathBind := binding.NewString()
	pathEntry := widget.NewEntryWithData(pathBind)
	pathEntry.SetPlaceHolder("Select a folder...")

	pathEntry.Validator = func(s string) error {
		if _, err := os.Stat(s); os.IsNotExist(err) {
			return errors.New("folder does not exist")
		}
		return nil
	}
	// pathLabel := widget.NewLabelWithData(pathBind)

	portBind := binding.NewInt()
	portEntry := widget.NewEntryWithData(binding.IntToString(portBind))
	portBind.Set(1000)
	portFormItem := widget.NewFormItem("Port", portEntry)

	jwtSecretBind := binding.NewString()
	jwtSecretEntry := widget.NewEntryWithData(jwtSecretBind)
	jwtSecretFormItem := widget.NewFormItem("Jwt Secret", jwtSecretEntry)

	jwtSecretEntry.Validator = func(s string) error {
		if len(s) < 5 {
			return errors.New("at least 5 characters long")
		}
		return nil
	}

	selectFolderDialog := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
		pathBind.Set(uri.Path())
	}, gi.MainWindow)

	buttonSelectFolder := widget.NewButton("Select folder", func() {
		selectFolderDialog.Show()
	})

	form := &widget.Form{
		BaseWidget: widget.BaseWidget{},
		Items: []*widget.FormItem{
			portFormItem,
			jwtSecretFormItem,
			widget.NewFormItem("Storage Folder", pathEntry),
			widget.NewFormItem("", buttonSelectFolder),
		},
		OnSubmit: func() {
			// Move to the main page.
			port, _ := portBind.Get()
			path, _ := pathBind.Get()
			jwtSecret, _ := jwtSecretBind.Get()

			config := &restapi.Config{
				FolderPath: path,
				JwtSecret:  jwtSecret,
				Port:       port,
			}

			// Save the config inside preferences.
			gi.ConfigStore.SaveConfig(config)

			// Set the config.
			restapi.GlobalConfig = config

			// Make the required files and folders
			err := MakeRequiredFiles(config)
			if err != nil {
				dialog.NewError(err, gi.MainWindow)
			}

			// Set the content to the other window.
			gi.SetContent(nextWindow)
		},
		OnCancel: func() {
			portBind.Set(1000)
			pathBind.Set("")
			jwtSecretBind.Set("")
		},
		SubmitText: "Save",
		CancelText: "Reset",
	}
	appTitle := widget.NewLabelWithStyle("mnimidamon server", fyne.TextAlignCenter, fyne.TextStyle{
		Bold:      true,
		Italic:    false,
		Monospace: true,
	})
	formContainer := container.NewPadded(form)
	c := container.NewBorder(appTitle, nil, nil, nil, formContainer)

	return &SetupWindow{
		Content:     c,
		Port:        portBind,
		FolderEntry: pathBind,
	}
}

func MakeRequiredFiles(config *restapi.Config) error {
	// Make file store folder
	if _, err := os.Stat(config.GetFileStoreFolderPath()); os.IsNotExist(err) {
		err := os.MkdirAll(config.GetFileStoreFolderPath(), 0700)
		if err != nil {
			return fmt.Errorf("could not create file store folder: %w", err)
		}
	}
	// Make database folder and file
	if _, err := os.Stat(config.GetDatabaseFilePath()); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(config.GetDatabaseFilePath()), 0700)
		if err != nil {
			return fmt.Errorf("could not create database folder: %w", err)
		}

		_, err = os.Create(config.GetDatabaseFilePath())
		if err != nil {
			return fmt.Errorf("could not create database file: %w", err)
		}
	}

	return nil
}
