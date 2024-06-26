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
	"runtime"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/spf13/cobra"

	"github.com/drewgonzales360/goenv/internal/pkg"
)

type GoEnvContextKey string

const (
	PermError string          = "you do not have access to %v"
	config    GoEnvContextKey = "config"

	// goVersionDarwinARMIntroduced is the version that the community started releasing binaries for
	// Apple Silicon
	goVersionDarwinARMIntroduced = "1.16"
)

// parseVersionArg ensures that the subcommands that accept a parameter only spcify one parameter.
func ValidateVersionArg(cmd *cobra.Command, args []string) error {
	if err := cobra.ExactArgs(1)(cmd, args); err != nil {
		return err
	}

	version, err := semver.NewVersion(args[0])
	if err != nil {
		return fmt.Errorf("invalid parameter: %w", err)
	}

	if !darwinArm(version) {
		return fmt.Errorf("go%s was not yet available for %s-%s", version, runtime.GOARCH, runtime.GOARCH)
	}

	return nil
}

// PostRun checks for new versions of Goenv and Go.
func PostRun(cmd *cobra.Command, _ []string) {
	v := semver.MustParse(cmd.Root().Version)

	if err := pkg.CheckLatestGoenv(cmd.Context(), v); err != nil {
		pkg.Debug(err.Error())
	}

	if err := pkg.CheckLatestGo(); err != nil {
		pkg.Debug(err.Error())
	}
}

// warnOnMissingPath does a best effort to let you know that Go can't be called.
func warnOnMissingPath(config *Config) {
	bin := config.GoenvRootDirectory + "/bin"
	if path := os.Getenv("PATH"); !strings.Contains(path, bin) {
		pkg.Info(fmt.Sprintf("%s is not in your PATH", bin))
		pkg.Info(fmt.Sprintf("export PATH=%s:$PATH", bin))
	}
}

func darwinArm(version *semver.Version) bool {
	return !(version.LessThan(semver.MustParse(goVersionDarwinARMIntroduced)) &&
		runtime.GOOS == "darwin" &&
		runtime.GOARCH == "arm64")
}
