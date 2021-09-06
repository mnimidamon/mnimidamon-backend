package main

import (
	"flag"
	"log"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/adapter/restapi/factory"
	"mnimidamonbackend/gui"
)

func main() {
	guiMode := flag.Bool("gui", true, "Run in GUI mode.")
	secret := flag.String("secret", "", "Secret for signing access keys.")
	port := flag.Int("port", 1000, "Port on which the server will listen to.")
	path := flag.String("path", "", "Path to file storage folder.")
	flag.Parse()

	if *guiMode {
		gsi, err := gui.NewGraphicalServerInterface()
		if err != nil {
			log.Fatalf("Error occured %v", err)
		}
		gsi.ShowAndRun()
	} else {

		if *path == "" || *secret == "" {
			log.Fatalf("Path and secret should be specified. Running in non gui mode requres port, path and secret flags set. More information with -help flag.")
		}

		restapi.GlobalConfig = &restapi.Config{
			FolderPath: *path,
			JwtSecret:  *secret,
			Port:       *port,
		}

		if err := gui.MakeRequiredFiles(restapi.GlobalConfig); err != nil {
			log.Fatalf("Could not make required files: %v", err)
		}

		s, err := factory.NewServer()

		if err != nil {
			log.Fatalf("Error when creating a new server: %v", err)
		}

		s.Port = restapi.GlobalConfig.Port

		if err := s.Serve(); err != nil {
			log.Fatalf("Error when starting the server: %v", err)
		}
	}
}