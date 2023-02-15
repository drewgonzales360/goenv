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
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/spf13/cobra"

	"github.com/drewgonzales360/goenv/internal/pkg"
)

type GoEnvContextKey string

const (
	PermError string          = "you do not have access to %v"
	config    GoEnvContextKey = "config"
)

// parseVersionArg ensures that the subcommands that accept a parameter
// only spcify one parameter.
func ValidateVersionArg(cmd *cobra.Command, args []string) error {
	if err := cobra.ExactArgs(1)(cmd, args); err != nil {
		return err
	}

	if _, err := semver.NewVersion(args[0]); err != nil {
		return fmt.Errorf("invalid parameter: %w", err)
	}

	return nil
}

// PostRunE checks for new versions of Goenv and Go.
func PostRunE(cmd *cobra.Command, _ []string) error {
	pkg.CheckLatestGoenv(cmd.Root().Version)
	pkg.CheckLatestGo()
	return nil
}

// warnOnMissingPath does a best effort to let you know that Go can't be called.
// This will sometimes warn unnecessarily if the user runs as root.
func warnOnMissingPath(config *Config) {
	bin := config.GoenvRootDirectory + "/bin"
	if path := os.Getenv("PATH"); !strings.Contains(path, bin) {
		pkg.Info(fmt.Sprintf("%s is not in your PATH", bin))
		pkg.Info(fmt.Sprintf("export PATH=%s:$PATH", bin))
	}
}
