package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/adapter/restapi/endpoints"
	"mnimidamonbackend/adapter/restapi/factory"
	"mnimidamonbackend/domain/constants"
	"mnimidamonbackend/gui/elements"
	"time"
)

type ServerRunningWindow struct {
	Content fyne.CanvasObject

	Server *endpoints.Server // Currently running server

	LogText *string
}

func (sr *ServerRunningWindow) StartANewServer() error {
	s, err := factory.NewServer()

	if err != nil {
		constants.Log("Error when creating a new server: %v", err)
		return err
	}

	s.Port = restapi.GlobalConfig.Port

	go func() {
		if err := s.Serve(); err != nil {
			constants.Log("Error when starting a new server: %v", err)
		}
	}()

	sr.Server = s
	return nil
}

func (sr *ServerRunningWindow) StopTheServer() error {
	if sr.Server == nil {
		constants.Log("There is no server running...")
		return nil
	}

	err := sr.Server.Shutdown()
	if err != nil {
		constants.Log("Error when shutting down server: %v", err)
	}

	sr.Server = nil
	return nil
}

func (sr *ServerRunningWindow) GetContent() fyne.CanvasObject {
	return sr.Content
}

func (sr *ServerRunningWindow) AppendToLogString(str string) {
	*sr.LogText = str + "\n" + *sr.LogText
}

func NewServerRunningWindow(gi *GraphicalInterface) *ServerRunningWindow {
	srw := &ServerRunningWindow{
		Content: nil,
		Server:  nil,
		LogText: new(string),
	}

	gi.MainWindow.Resize(fyne.Size{
		Width:  500,
		Height: 250,
	})

	toolbarContainer := container.NewMax()
	var toolbarStart, toolbarStop *widget.Toolbar
	var startServerAction, stopServerAction widget.ToolbarItem

	startServerAction = widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
		err := srw.StartANewServer()
		if err != nil {
			return
		}
		// Replace the toolbar.
		toolbarContainer.Remove(toolbarStart)
		toolbarContainer.Add(toolbarStop)
	})

	stopServerAction = widget.NewToolbarAction(theme.MediaStopIcon(), func() {
		err := srw.StopTheServer()
		if err != nil {
			return
		}
		// Replace the toolbar.
		toolbarContainer.Remove(toolbarStop)
		toolbarContainer.Add(toolbarStart)
	})

	resetAppSettingsAction := widget.NewToolbarAction(theme.SettingsIcon(), func() {
		d := dialog.NewConfirm("Are you sure?", "This will delete your server settings.\n These will be reconfigurable on next server startup.", func(b bool) {
			if b {
				gi.ConfigStore.DeleteConfig()
				gi.MainWindow.Close()
			}
		}, gi.MainWindow)
		d.Show()
	})

	toolbarLabel := elements.NewToolbarLabel("mnimidamon server")

	toolbarStart = widget.NewToolbar(
		resetAppSettingsAction,
		toolbarLabel,
		widget.NewToolbarSpacer(),
		startServerAction,
	)

	toolbarStop = widget.NewToolbar(
		resetAppSettingsAction,
		toolbarLabel,
		widget.NewToolbarSpacer(),
		stopServerAction,
	)

	toolbarContainer.Add(toolbarStart)

	logField := widget.NewMultiLineEntry()
	logField.SetPlaceHolder("Start the server when ready...")
	logField.Disable()

	// Replace the global logging to output the log inside the gui.
	logFn := func(s string, i ...interface{}) {
		srw.AppendToLogString(fmt.Sprintf("%v %v", time.Now().Format("02-01-2006 15:04:05 "), fmt.Sprintf(s, i...)))
		logField.SetText(*srw.LogText)
	}

	constants.Log = logFn

	mainContent := canvas.NewRectangle(theme.BackgroundColor())
	mainContent.SetMinSize(fyne.NewSize(400, 4))

	srw.Content = container.NewBorder(toolbarContainer, nil, nil, nil, mainContent, logField)
	return srw
}
