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
package pkg_test

import (
	"path"
	"testing"

	"github.com/Masterminds/semver"
	"github.com/stretchr/testify/assert"

	"github.com/drewgonzales360/goenv/internal/pkg"
)

const testDir = "/tmp/goenv/"

func TestDownloadAndUntar(t *testing.T) {

	testCases := []string{
		"1.18",
		"1.20.5",
		"1.21",
	}

	for _, version := range testCases {
		goVersion := semver.MustParse(version)
		tarballPath, err := pkg.DownloadFile(goVersion)
		assert.NoError(t, err)

		installDir := path.Join(testDir, goVersion.String())
		err = pkg.ExtractTarGz(tarballPath, installDir)
		assert.NoError(t, err)

		assert.DirExists(t, installDir)
	}
}
