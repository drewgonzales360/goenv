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
	"os"
	"path"

	"github.com/Masterminds/semver/v3"
	"github.com/spf13/cobra"

	"github.com/drewgonzales360/goenv/internal/pkg"
)

func UninstallCommand(cmd *cobra.Command, args []string) error {
	version := args[0]
	config := pkg.ReadConfig()

	if err := uninstall(config, version); err != nil {
		return err
	}

	return nil
}

func uninstall(config *pkg.Config, version string) error {
	if inaccessible := pkg.CheckRW(config.GoenvRootDirectory, config.GoenvInstallDirectory); len(inaccessible) > 0 {
		return fmt.Errorf(PermError, inaccessible)
	}

	goVersion, err := semver.NewVersion(version)
	if err != nil {
		return fmt.Errorf("could not parse %s: %w", version, err)
	}

	if err := pkg.CheckInstalled(config.GoenvInstallDirectory, goVersion.String()); err != nil {
		return fmt.Errorf("go version not installed: %w", err)
	}

	if err = os.RemoveAll(path.Join(config.GoenvInstallDirectory, goVersion.String())); err != nil {
		return fmt.Errorf("could not uninstall go: %w", err)
	}

	pkg.Success("Uninstalled Go " + goVersion.String())

	// Use another version
	versions, err := os.ReadDir(config.GoenvInstallDirectory)
	if err != nil {
		return err
	}
	if len(versions) == 0 {
		pkg.Warn("No other Go versions installed")
		return nil
	}

	latestVersion := versions[len(versions)-1].Name()
	if err := use(config, latestVersion); err != nil {
		return fmt.Errorf("could not default to %s: %s", latestVersion, err)
	}

	return nil
}
