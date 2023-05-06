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
package pkg

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"runtime"
	"sort"

	"github.com/Masterminds/semver"
	"github.com/fatih/color"
	"github.com/google/go-github/v48/github"
)

const (
	installInstructions = "  curl -sSL 'https://github.com/drewgonzales360/goenv/releases/download/%s/goenv-%s-%s-%s.tar.gz' | sudo tar -xzv -C /usr/local/bin\n"
	releaseDataPath     = "/var/tmp/goenv-latest.txt"
)

// ReleaseCollection implements the necessary interface to sort. Len, Less, and Swap
type ReleaseCollection []*github.RepositoryRelease

func (c ReleaseCollection) Len() int {
	return len(c)
}

func (c ReleaseCollection) Less(i, j int) bool {
	return c[i].PublishedAt.Before(c[j].PublishedAt.Time)
}

func (c ReleaseCollection) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c ReleaseCollection) latest() *github.RepositoryRelease {
	return c[len(c)-1]
}

// CheckLatestGoenv checks the Github releases for a new version of Goenv. If one is found, the
// one-line install instructions are printed to the console.
func CheckLatestGoenv(currentVersion *semver.Version) error {
	gh := github.NewClient(nil)
	repoRelease, _, err := gh.Repositories.ListReleases(context.Background(), "drewgonzales360", "goenv", nil)
	if err != nil {
		return err
	}

	if len(repoRelease) < 1 {
		return fmt.Errorf("found no releases for drewgonzales360/goenv")
	}

	releaseCollection := ReleaseCollection(repoRelease)
	sort.Sort(releaseCollection)
	latestRelease := releaseCollection.latest()

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

	newReleaseData := []byte(fmt.Sprintf("%+v\n", releases))
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

	newMsg := "A new version of Go is available:"
	if len(releases) > 1 {
		newMsg = "New versions of Go are available:"
	}

	if !bytes.Equal(oldReleaseData, newReleaseData) {
		color.Green(newMsg)
		gvl := CreateGoVersionList(releases)
		Print(gvl, "  ")
	}

	return nil
}
