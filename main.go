package main

import (
	"log"
	"os"

	"github.com/drewgonzales360/alfred/cmd"
	"github.com/drewgonzales360/alfred/version"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    version.AppName,
		Usage:   "tempalate for go command line tools",
		Version: version.Version,
		Commands: []*cli.Command{
			{
				Name:   "run",
				Usage:  "run the server",
				Action: cmd.Run,
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
