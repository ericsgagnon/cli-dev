package main

import (
	"cli-test/cmd"
	"log"
	"os"
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
