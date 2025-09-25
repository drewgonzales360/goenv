package cmd

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
import (
	"fmt"

	"github.com/drewgonzales360/goenv/internal/pkg"
	"github.com/spf13/cobra"
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
)

func ConfigCommand(cmd *cobra.Command, _ []string) {
	config := pkg.ReadConfig()
	fmt.Print(config)
	config.WarnOnMissingPath()
}
