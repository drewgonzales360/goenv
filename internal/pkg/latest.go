package pkg

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
	"bytes"
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/Masterminds/semver/v3"
	"github.com/fatih/color"
	"github.com/google/go-github/v74/github"
)

const (
	installInstructions = "  curl -sSL 'https://github.com/drewgonzales360/goenv/releases/download/v%s/goenv-%s-%s-v%s.tar.gz' | sudo tar -xzv -C /usr/local/bin\n"
	releaseDataPath     = "/var/tmp/goenv-latest.txt"
)

// CheckLatestGoenv checks the Github releases for a new version of Goenv. If one is found, the
// one-line install instructions are printed to the console.
func CheckLatestGoenv(ctx context.Context, currentVersion *semver.Version) error {
	gh := github.NewClient(nil)
	latestRelease, _, err := gh.Repositories.GetLatestRelease(ctx, "drewgonzales360", "goenv")
	if err != nil {
		return err
	}

	latest, err := semver.NewVersion(latestRelease.GetTagName())
	if err != nil {
		return err
	}

	if currentVersion.LessThan(latest) {
		color.Green(fmt.Sprintf("\nA new version of goenv has been released. %s ➡️ %s", currentVersion, latest))
		fmt.Printf(installInstructions, latest, runtime.GOOS, runtime.GOARCH, latest)
	}

	return nil
}

// CheckLatestGo looks for new stable versions of Go. If new stable versions have been released
// since the last check, then we'll let the user know. We only print this message once per new set
// of releases.
func CheckLatestGo() error {
	releases, err := ListAvailableVersions(false)
	if err != nil {
		return err
	}

	newReleaseData := fmt.Appendf(nil, "%+v\n", releases)
	defer func() {
		err := os.WriteFile(releaseDataPath, newReleaseData, 0666)
		if err != nil {
			Debug(err.Error())
		}
	}()

	oldReleaseData, err := os.ReadFile(releaseDataPath)
	if err != nil {
		oldReleaseData = newReleaseData
		Debug(err.Error())
	}

	if !bytes.Equal(oldReleaseData, newReleaseData) {
		color.Green("New versions of Go are available:")
		gvl := CreateGoVersionList(releases)
		Print(gvl, "  ")
	}

	return nil
}
