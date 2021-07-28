package main

import (
	"log"
	"mnimidamonbackend/gui"
)

func main() {
	gsi, err := gui.NewGraphicalServerInterface()
	if err != nil {
		log.Fatalf("Error occured %v", err)
	}
	gsi.ShowAndRun()
}
