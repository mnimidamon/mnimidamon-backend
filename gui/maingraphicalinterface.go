package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/gui/appicons"
)

type GraphicalInterface struct {
	App        fyne.App
	MainWindow fyne.Window // Main application window.

	ConfigStore *ConfigStore

	SetupWindow         *SetupWindow
	ServerRunningWindow *ServerRunningWindow
}

func (gui GraphicalInterface) SetContent(cg ContentGetter) {
	gui.MainWindow.SetContent(cg.GetContent())
}

func (gui *GraphicalInterface) ShowAndRun() {
	// Are the setting already saved?
	conf := gui.ConfigStore.GetConfig()
	if conf == nil {
		// Display the setup window as we need to get the config data.
		gui.SetContent(gui.SetupWindow)
	} else {
		// Save the global config
		restapi.GlobalConfig = conf
		// Display the server running window.
		gui.SetContent(gui.ServerRunningWindow)
	}

	gui.MainWindow.ShowAndRun()
}

func NewGraphicalServerInterface() (*GraphicalInterface, error) {
	mdApp := app.NewWithID("server.mnimidamon.marmiha.com")
	mdApp.SetIcon(appicons.ResourceMnimidamonServerIconPng)

	mainWindow := mdApp.NewWindow(" ")

	gi := &GraphicalInterface{
		App:        mdApp,
		MainWindow: mainWindow,

		ConfigStore: &ConfigStore{
			Pref: mdApp.Preferences(),
		},

		SetupWindow:         nil,
		ServerRunningWindow: nil,
	}

	serverRunningWindow := NewServerRunningWindow(gi)

	// Create the setup window.
	// Set the callback when the fields are entered.
	setupWindow := NewSetupWindow(gi, serverRunningWindow)

	gi.SetupWindow = setupWindow
	gi.ServerRunningWindow = serverRunningWindow

	return gi, nil
}
