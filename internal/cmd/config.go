// ///////////////////////////////////////////////////////////////////////
// Copyright 2024 Drew Gonzales
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
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	// The file placed at GoenvRootDirectory is a symlink to the GoenvInstallDirectory
	GoenvRootDirectory string

	// Install directory defaults to /usr/local/goenv and can be configured with
	GoenvInstallDirectory string
}

const (
	DefaultGoenvRootDirectory = "/usr/local/go"
	DefaultGoInstallDirectory = "/usr/local/goenv"

	GoEnvRootDirEnvVar    = "GOENV_ROOT_DIR"
	GoEnvInstallDirEnvVar = "GOENV_INSTALL_DIR"

	// Used to let the user know how a config variable was set.
	configSetByDefault = "\t(default)\n"
	configSetByEnv     = "\t(set by environment variable)\n"
)

func ConfigCommand(cmd *cobra.Command, _ []string)  {
	config := ReadConfig()
	fmt.Print(config)
	warnOnMissingPath(config)
}

// ReadConfig reads the environment variables for a user and creates a config. If we need any
// additional config, it'll be parsed in here.
func ReadConfig() *Config {
	return &Config{
		GoenvRootDirectory:    viper.GetString(GoEnvRootDirEnvVar),
		GoenvInstallDirectory: viper.GetString(GoEnvInstallDirEnvVar),
	}
}

// String prints the config to a tabwriter so that the columns are aligned when it's fmt.Print'ed to
// the terminal.
func (c *Config) String() string {
	buf := bytes.Buffer{}
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', 0)

	rootDir := fmt.Sprintf("%s:\t%s", GoEnvRootDirEnvVar, c.GoenvRootDirectory)
	if c.GoenvRootDirectory == DefaultGoenvRootDirectory {
		rootDir += configSetByDefault
	} else {
		rootDir += configSetByEnv
	}
	w.Write([]byte(rootDir))

	installDir := fmt.Sprintf("%s:\t%s", GoEnvInstallDirEnvVar, c.GoenvInstallDirectory)
	if c.GoenvInstallDirectory == DefaultGoInstallDirectory {
		installDir += configSetByDefault
	} else {
		installDir += configSetByEnv
	}
	w.Write([]byte(installDir))

	w.Flush()

	return buf.String()
}
