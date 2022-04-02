package main

import (
	"os"

	"github.com/drewgonzales360/goenv/cmd"
	"github.com/drewgonzales360/goenv/pkg"
	"github.com/drewgonzales360/goenv/version"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    version.AppName,
		Usage:   "Manages multiple go versions for linux",
		Version: version.Version().Original(),
		Commands: []*cli.Command{
			{
				Name:    "install",
				Usage:   "install a go version",
				Aliases: []string{"i"},
				Action:  cmd.InstallCommand,
			},
			{
				Name:    "uninstall",
				Usage:   "uninstall a go version",
				Aliases: []string{"rm"},
				Action:  cmd.UninstallCommand,
			},
			{
				Name:    "use",
				Usage:   "use a go version",
				Aliases: []string{"u"},
				Action:  cmd.UseCommand,
			},
			{
				Name:    "list",
				Usage:   "list all available go versions",
				Aliases: []string{"ls", "l"},
				Action:  cmd.ListCommand,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		pkg.Fail(err.Error())
		os.Exit(1)
	}
}
