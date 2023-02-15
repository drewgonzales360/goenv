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
	"path"

	"github.com/Masterminds/semver"
	"github.com/spf13/cobra"

	"github.com/drewgonzales360/goenv/internal/pkg"
)

func UninstallCommand(cmd *cobra.Command, args []string) error {
	version := args[0]
	config := ReadConfig()

	if err := uninstall(config, version); err != nil {
		return err
	}

	return nil
}

func uninstall(config *Config, versionToDelete string) error {
	if inaccessible := pkg.CheckRW(config.GoenvRootDirectory, config.GoenvInstallDirectory); len(inaccessible) > 0 {
		return fmt.Errorf(PermError, inaccessible)
	}

	// Remove the old Go version
	goVersion, err := semver.NewVersion(versionToDelete)
	if err != nil {
		return fmt.Errorf("could not parse %s: %w", versionToDelete, err)
	}

	if err = os.RemoveAll(path.Join(config.GoenvInstallDirectory, goVersion.Original())); err != nil {
		return fmt.Errorf("could not uninstall go: %w", err)
	}

	pkg.Success("Uninstalled Go " + goVersion.Original())

	// Use another version
	versions, err := os.ReadDir(config.GoenvInstallDirectory)
	if err != nil {
		return err
	}
	if len(versions) == 0 {
		pkg.Warn("no other installed go versions")
		return nil
	}

	defaultVersion := versions[len(versions)-1].Name()
	if err := use(config, defaultVersion); err == nil {
		pkg.Debug(fmt.Sprintf("could not default to %s: %s", defaultVersion, err))
	}

	return nil
}
