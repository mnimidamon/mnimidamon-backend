package gui

import (
	"fyne.io/fyne/v2"
	"mnimidamonbackend/adapter/restapi"
)

var (
	IsStoredPrefKey  = "is_config_stored"
	PortPrefKey      = "port"
	FolderPrefKey    = "folder"
	JwtSecretPrefKey = "jwt_secret_key"
)

type ConfigStore struct {
	Pref fyne.Preferences
}

func (cs ConfigStore) DeleteConfig() {
	cs.Pref.SetBool(IsStoredPrefKey, false)
}

func (cs ConfigStore) GetConfig() *restapi.Config {
	if !cs.Pref.Bool(IsStoredPrefKey) {
		return nil
	}

	port := cs.Pref.Int(PortPrefKey)
	folder := cs.Pref.String(FolderPrefKey)
	jwtSecret := cs.Pref.String(JwtSecretPrefKey)

	return &restapi.Config{
		FolderPath: folder,
		JwtSecret:  jwtSecret,
		Port:       port,
	}
}

func (cs ConfigStore) SaveConfig(c *restapi.Config) {
	cs.Pref.SetInt(PortPrefKey, c.Port)
	cs.Pref.SetString(JwtSecretPrefKey, c.JwtSecret)
	cs.Pref.SetString(FolderPrefKey, c.FolderPath)

	cs.Pref.SetBool(IsStoredPrefKey, true)
}
