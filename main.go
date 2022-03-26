package main

import (
	"log"
	"os"

	"github.com/drewgonzales360/goenv/cmd"
	"github.com/drewgonzales360/goenv/version"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    version.AppName,
		Usage:   "Manages multiple go versions for linux",
		Version: version.Version,
		Commands: []*cli.Command{
			{
				Name:   "install",
				Usage:  "install a go version",
				Action: cmd.Install,
			},
			{
				Name:   "uninstall",
				Usage:  "uninstall a go version",
				Action: cmd.Uninstall,
			},
			{
				Name:   "use",
				Usage:  "use a go version",
				Action: cmd.Use,
			},
			{
				Name:   "list",
				Usage:  "list all available go versions",
				Action: cmd.List,
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
