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
	"fmt"
	"runtime"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
)

// go1_20_5Checksums are the sha256 sums of the go1.20.5 archives, keyed by GOOS-GOARCH.
var go1_20_5Checksums = map[string]string{
	"darwin-arm64": "94ad76b7e1593bb59df7fd35a738194643d6eed26a4181c94e3ee91381e40459",
	"linux-amd64":  "d7ec48cde0d3d2be2c69203bc3e0a44de8660b9c09a6e85c4732a3f7dc442612",
}

func TestPrintInstallScript(t *testing.T) {
	platform := fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
	checksum, ok := go1_20_5Checksums[platform]
	if !ok {
		t.Skipf("no known go1.20.5 checksum for %s", platform)
	}

	tests := []struct {
		name       string
		rootDir    string
		installDir string
		version    *semver.Version
		want       string
	}{
		{
			name:       "Test PrintInstallScript",
			rootDir:    "/usr/local/go",
			installDir: "/usr/local/goenv/1.20.5",
			version:    semver.MustParse("1.20.5"),
			want:       `curl -o /tmp/go1.20.5.%[1]s.tar.gz -sSL https://go.dev/dl/go1.20.5.%[1]s.tar.gz
echo "%[2]s /tmp/go1.20.5.%[1]s.tar.gz" | sha256sum -c -
mkdir -p /usr/local/goenv/1.20.5
tar --strip-components=1 -xzf /tmp/go1.20.5.%[1]s.tar.gz -C /usr/local/goenv/1.20.5
ln -s /usr/local/goenv/1.20.5 /usr/local/go
export PATH=/usr/local/go/bin:$PATH`,
		},
		{
			name:       "global",
			rootDir:    "/usr/local/go",
			installDir: "/usr/local/go",
			version:    semver.MustParse("1.20.5"),
			want:       `curl -o /tmp/go1.20.5.%[1]s.tar.gz -sSL https://go.dev/dl/go1.20.5.%[1]s.tar.gz
echo "%[2]s /tmp/go1.20.5.%[1]s.tar.gz" | sha256sum -c -
mkdir -p /usr/local/go
tar --strip-components=1 -xzf /tmp/go1.20.5.%[1]s.tar.gz -C /usr/local/go
export PATH=/usr/local/go/bin:$PATH`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := InstallScript(tt.rootDir, tt.installDir, tt.version)
			assert.Equal(t, fmt.Sprintf(tt.want, platform, checksum), actual)
		})
	}
}
