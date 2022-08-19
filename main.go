package main

import (
	"fmt"
	"os"

	"github.com/Masterminds/semver"
	"github.com/urfave/cli/v2"

	"github.com/drewgonzales360/goenv/internal/cmd"
	"github.com/drewgonzales360/goenv/internal/pkg"
)

const appName = "goenv"

// Semver is set in the Makefile
var Semver = "unknown; please create an issue for the maintainers"

func main() {
	app := &cli.App{
		Name: appName,
		Authors: []*cli.Author{
			{
				Name:  "Drew Gonzales",
				Email: "github.com/drewgonzales360",
			},
		},
		Version:   semver.MustParse(Semver).Original(),
		Usage:     "Manages multiple go versions for linux. See https://go.dev/dl for available versions.",
		UsageText: fmt.Sprintf("%s <command> [version]", appName),
		Commands: []*cli.Command{
			{
				Name:      "install",
				Usage:     "Install a Go version. Usually in the form 1.18, 1.9, 1.17.8.",
				UsageText: fmt.Sprintf("%s install [version]", appName),
				Aliases:   []string{"i"},
				Before:    cmd.BeforeActionParseConfig,
				Action:    cmd.InstallCommand,
			},
			{
				Name:      "uninstall",
				Usage:     "Uninstall a Go version",
				UsageText: fmt.Sprintf("%s uninstall [version]", appName),
				Aliases:   []string{"rm"},
				Before:    cmd.BeforeActionParseConfig,
				Action:    cmd.UninstallCommand,
			},
			{
				Name:      "use",
				Usage:     "Use a Go version",
				UsageText: fmt.Sprintf("%s use [version]", appName),
				Aliases:   []string{"u"},
				Before:    cmd.BeforeActionParseConfig,
				Action:    cmd.UseCommand,
			},
			{
				Name:    "list",
				Usage:   "List all available Go versions",
				Aliases: []string{"ls", "l"},
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "stable",
						Aliases: []string{"s"},
					},
					&cli.BoolFlag{
						Name:    "all",
						Aliases: []string{"a"},
					},
				},
				Before: cmd.BeforeActionParseConfig,
				Action: cmd.ListCommand,
			},
			{
				Name:    "config",
				Usage:   "Prints the current config.",
				Aliases: []string{"c"},
				Before:  cmd.BeforeActionParseConfig,
				Action:  cmd.ConfigCommand,
			},
		},
		HideHelpCommand: true,
	}

	if err := app.Run(os.Args); err != nil {
		pkg.Error(err.Error())
		os.Exit(1)
	}
}
