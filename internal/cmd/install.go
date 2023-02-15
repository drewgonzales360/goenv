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
	"path"

	"github.com/Masterminds/semver"
	"github.com/spf13/cobra"

	"github.com/drewgonzales360/goenv/internal/pkg"
)

func InstallCommand(cmd *cobra.Command, args []string) error {
	version := args[0]
	config := ReadConfig()

	if err := install(config, version); err != nil {
		return err
	}

	return nil
}

func install(config *Config, version string) error {
	if inaccessible := pkg.CheckRW(config.GoenvRootDirectory, config.GoenvInstallDirectory); len(inaccessible) > 0 {
		return fmt.Errorf(PermError, inaccessible)
	}

	goVersion, err := semver.NewVersion(version)
	if err != nil {
		return fmt.Errorf("could not parse %s: %w", version, err)
	}

	filePath, err := pkg.DownloadFile(goVersion)
	if err != nil {
		return fmt.Errorf("could not download go: %w", err)
	}

	err = pkg.ExtractTarGz(filePath, path.Join(config.GoenvInstallDirectory, goVersion.Original()))
	if err != nil {
		return fmt.Errorf("could not extract go: %w", err)
	}

	err = use(config, version)
	if err != nil {
		return err
	}

	return nil
}
