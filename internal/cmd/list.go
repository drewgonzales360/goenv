// ///////////////////////////////////////////////////////////////////////
// Copyright 2022 Drew Gonzales
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
package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"

	"github.com/drewgonzales360/goenv/internal/pkg"
)

func ListCommand(c *cli.Context) error {
	config, err := parseConfig(c)
	if err != nil {
		return err
	}

	versions, err := os.ReadDir(config.GoenvInstallDirectory)
	if err != nil {
		return err
	}

	names := []string{}
	for _, version := range versions {
		names = append(names, version.Name())
	}
	installed := pkg.CreateGoVersionList(names)
	color.New(color.FgCyan, color.Bold).Println("Installed Versions:")
	pkg.Print(installed, "")

	all := c.Bool("all")
	printAvailable := c.Bool("stable") || all
	if printAvailable {
		color.New(color.FgCyan, color.Bold).Println("Available Versions:")
		versions, err := pkg.ListAvailableVersions(all)
		if err != nil {
			return err
		}

		gvl := pkg.CreateGoVersionList(versions)
		pkg.Print(gvl, "")
	}

	return nil
}
