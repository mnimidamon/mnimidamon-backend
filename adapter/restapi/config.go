package restapi

import (
	"path/filepath"
)

// Read by the mnimidamon configure function to create the right objects,
// as we can not pass the information to the function.
var GlobalConfig *Config

var (
	fileStoreFolder = "filestore"
	databaseFolder  = "database"
	databaseFile    = "mnimidamon.db"
)

type Config struct {
	FolderPath string // Base folder path.
	JwtSecret  string // JwtAuthentication secret
	Port       int    // The port on which the server will listen to.
}

func (c Config) GetDatabaseFilePath() string {
	return filepath.Join(c.FolderPath, databaseFolder, databaseFile)
}

func (c Config) GetFileStoreFolderPath() string {
	return filepath.Join(c.FolderPath, fileStoreFolder)
}
