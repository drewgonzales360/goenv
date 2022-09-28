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
		Usage:     "Manages multiple Go versions. See https://go.dev/dl for available versions or `goenv ls -s`.",
		UsageText: fmt.Sprintf("%s <command> [version]", appName),
		Commands: []*cli.Command{
			{
				Name:      "install",
				Usage:     "Install a Go version. Usually in the form 1.18, 1.9, 1.17.8.",
				UsageText: fmt.Sprintf("ex: %s install 1.19.1", appName),
				Aliases:   []string{"i"},
				Before:    cmd.BeforeActionParseConfig,
				Action:    cmd.InstallCommand,
			},
			{
				Name:      "uninstall",
				Usage:     "Uninstall a Go version.",
				UsageText: fmt.Sprintf("ex: %s uninstall 1.17", appName),
				Aliases:   []string{"rm"},
				Before:    cmd.BeforeActionParseConfig,
				Action:    cmd.UninstallCommand,
			},
			{
				Name:      "use",
				Usage:     "Use a Go version.",
				UsageText: fmt.Sprintf("ex: %s use 1.18", appName),
				Aliases:   []string{"u"},
				Before:    cmd.BeforeActionParseConfig,
				Action:    cmd.UseCommand,
			},
			{
				Name:    "list",
				Usage:   "List available Go versions.",
				Aliases: []string{"ls", "l"},
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "stable",
						Aliases: []string{"s"},
						Usage:   "lists all actively maintained versions of Go",
					},
					&cli.BoolFlag{
						Name:    "all",
						Aliases: []string{"a"},
						Usage:   "lists all versions of Go",
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
		HideHelpCommand:        true,
		UseShortOptionHandling: true,
		Suggest:                true,
	}

	if err := app.Run(os.Args); err != nil {
		pkg.Error(err.Error())
		os.Exit(1)
	}
}
