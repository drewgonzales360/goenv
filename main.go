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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/drewgonzales360/goenv/internal/cmd"
)

const appName = "goenv"

// Semver is set in the Makefile
var Semver = "unknown; please create an issue for the maintainers"

func main() {
	viper.BindEnv(cmd.GoEnvRootDirEnvVar)
	viper.SetDefault(cmd.GoEnvRootDirEnvVar, cmd.DefaultGoenvRootDirectory)
	viper.BindEnv(cmd.GoEnvInstallDirEnvVar)
	viper.SetDefault(cmd.GoEnvInstallDirEnvVar, cmd.DefaultGoInstallDirectory)

	installCmd := &cobra.Command{
		Use:     "install",
		Short:   "Install a Go version. Usually in the form 1.18, 1.9, 1.17.8.",
		Example: fmt.Sprintf("ex: %s install 1.19.1", appName),
		Aliases: []string{"i"},
		RunE:    cmd.InstallCommand,
		Args:    cmd.ValidateVersionArg,
		PostRun: cmd.PostRun,
	}

	uninstallCmd := &cobra.Command{
		Use:     "uninstall",
		Short:   "Uninstall a Go version.",
		Example: fmt.Sprintf("ex: %s uninstall 1.17", appName),
		Aliases: []string{"rm"},
		RunE:    cmd.UninstallCommand,
		Args:    cmd.ValidateVersionArg,
		PostRun: cmd.PostRun,
	}

	useCmd := &cobra.Command{
		Use:     "use",
		Short:   "Switch the current Go version to use whichever version in specified and installed.",
		Example: fmt.Sprintf("ex: %s use 1.18", appName),
		Aliases: []string{"u"},
		RunE:    cmd.UseCommand,
		Args:    cmd.ValidateVersionArg,
		PostRun: cmd.PostRun,
	}

	listCmd := &cobra.Command{
		Use:     "list",
		Short:   "List all installed available Go versions.",
		Aliases: []string{"ls", "l"},
		RunE:    cmd.ListCommand,
		Args:    cobra.ExactArgs(0),
		PostRun: cmd.PostRun,
	}
	listCmd.Flags().BoolP("stable", "s", false, "Print out only new stable releases.")
	listCmd.Flags().BoolP("all", "a", false, "Print out all releases.")

	configCmd := &cobra.Command{
		Use:     "config",
		Short:   "Print out the current config",
		Aliases: []string{"c"},
		RunE:    cmd.ConfigCommand,
		Args:    cobra.ExactArgs(0),
		PostRun: cmd.PostRun,
	}

	rootCmd := &cobra.Command{Version: Semver}
	rootCmd.AddCommand(installCmd, uninstallCmd, useCmd, listCmd, configCmd)

	rootCmd.Execute()
}
