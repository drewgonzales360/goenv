// ///////////////////////////////////////////////////////////////////////
// Copyright 2023 Drew Gonzales
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ///////////////////////////////////////////////////////////////////////
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
				After:     cmd.AfterAction,
				Action:    cmd.InstallCommand,
			},
			{
				Name:      "uninstall",
				Usage:     "Uninstall a Go version.",
				UsageText: fmt.Sprintf("ex: %s uninstall 1.17", appName),
				Aliases:   []string{"rm"},
				Before:    cmd.BeforeActionParseConfig,
				After:     cmd.AfterAction,
				Action:    cmd.UninstallCommand,
			},
			{
				Name:      "use",
				Usage:     "Use a Go version.",
				UsageText: fmt.Sprintf("ex: %s use 1.18", appName),
				Aliases:   []string{"u"},
				Before:    cmd.BeforeActionParseConfig,
				After:     cmd.AfterAction,
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
				After:  cmd.AfterAction,
				Action: cmd.ListCommand,
			},
			{
				Name:    "config",
				Usage:   "Prints the current config.",
				Aliases: []string{"c"},
				Before:  cmd.BeforeActionParseConfig,
				After:   cmd.AfterAction,
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
