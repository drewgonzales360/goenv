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
	"os/exec"
	"path"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/spf13/cobra"

	"github.com/drewgonzales360/goenv/internal/pkg"
)

func UseCommand(cmd *cobra.Command, args []string) error {
	version := args[0]
	config := ReadConfig()

	if err := use(config, version); err != nil {
		return err
	}

	return nil
}

func use(config *Config, version string) error {
	if inaccessible := pkg.CheckRW(config.GoenvRootDirectory, config.GoenvInstallDirectory); len(inaccessible) > 0 {
		return fmt.Errorf(PermError, inaccessible)
	}

	goVersion, err := semver.NewVersion(version)
	if err != nil {
		return fmt.Errorf("could not parse %s: %w", version, err)
	}

	if err = link(config, goVersion); err != nil {
		return err
	}

	goCmd := "go"
	if isRoot() {
		// The root user has no idea what the is in the path of the normal user, so we call the
		// binary directly.
		goCmd = path.Join(config.GoenvRootDirectory, "bin", goCmd)
	} else {
		// If the non-root user is calling, then we should check if it's in their path and warn if
		// it isn't.
		warnOnMissingPath(config)
	}
	output, err := exec.Command(goCmd, "version").Output()
	if err != nil {
		return err
	}

	pkg.Success(fmt.Sprintf("Using %s", strings.TrimSuffix(string(output), "\n")))
	return nil
}

func link(config *Config, goVersion *semver.Version) error {
	// Remove the old symlink
	if _, err := os.Stat(config.GoenvRootDirectory); err == nil {
		if err = os.Remove(config.GoenvRootDirectory); err != nil {
			return fmt.Errorf("could not remove %s: %w", config.GoenvRootDirectory, err)
		}
	}

	gvs := goVersion.String()
	goInstallation := path.Join(config.GoenvInstallDirectory, gvs)
	if _, err := os.Stat(goInstallation); err != nil {
		pkg.Debug(err.Error())
		return fmt.Errorf("could not find go version %s. goenv install %s", gvs, gvs)
	}

	if err := os.Symlink(goInstallation, config.GoenvRootDirectory); err != nil {
		return fmt.Errorf("could not link: %w", err)
	}

	return nil
}
func isRoot() bool {
	return os.Geteuid() == 0
}
