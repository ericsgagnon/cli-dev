package main

import (
	"log"
	"os"

	"github.com/ericsgagnon/cli-dev/cmd"
)

func main() {
	app, err := cmd.NewApp()
	if err != nil {
		log.Fatalln(err)
	}
	err = (app).Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(0)
}
