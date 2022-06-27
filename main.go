package main

import (
	"fmt"
	"os"

	"github.com/Masterminds/semver"
	"github.com/drewgonzales360/goenv/internal/cmd"
	"github.com/drewgonzales360/goenv/internal/pkg"
	"github.com/urfave/cli/v2"
)

const AppName = "goenv"

// Version semvers the app
var Semver string = "unknown; please create an issue for the maintainers"

func main() {
	app := &cli.App{
		Name: AppName,
		Authors: []*cli.Author{
			{
				Name:  "Drew Gonzales",
				Email: "github.com/drewgonzales360",
			},
		},
		Version:   semver.MustParse(Semver).Original(),
		Usage:     "Manages multiple go versions for linux. See https://go.dev/dl for available versions.",
		UsageText: fmt.Sprintf("%s <command> [version]", AppName),
		Commands: []*cli.Command{
			{
				Name:      "install",
				Usage:     "Install a Go version. Usually in the form 1.18, 1.9, 1.17.8.",
				UsageText: fmt.Sprintf("%s install [version]", AppName),
				Aliases:   []string{"i"},
				Before:    cmd.BeforeActionParseConfig,
				Action:    cmd.InstallCommand,
			},
			{
				Name:      "uninstall",
				Usage:     "Uninstall a go version",
				UsageText: fmt.Sprintf("%s uninstall [version]", AppName),
				Aliases:   []string{"rm"},
				Before:    cmd.BeforeActionParseConfig,
				Action:    cmd.UninstallCommand,
			},
			{
				Name:      "use",
				Usage:     "Use a go version",
				UsageText: fmt.Sprintf("%s use [version]", AppName),
				Aliases:   []string{"u"},
				Before:    cmd.BeforeActionParseConfig,
				Action:    cmd.UseCommand,
			},
			{
				Name:    "list",
				Usage:   "List all available go versions",
				Aliases: []string{"ls", "l"},
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "available",
						Aliases: []string{"a"},
					},
				},
				Before: cmd.BeforeActionParseConfig,
				Action: cmd.ListCommand,
			},
			{
				Name:    "config",
				Usage:   "prints the current config",
				Aliases: []string{"c"},
				Before:  cmd.BeforeActionParseConfig,
				Action:  cmd.ConfigCommand,
			},
		},
		HideHelpCommand: true,
	}

	if err := app.Run(os.Args); err != nil {
		pkg.Fail(err.Error())
		os.Exit(1)
	}
}
