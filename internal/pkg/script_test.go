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
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
)

func TestPrintInstallScript(t *testing.T) {
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
			want:       `curl -o /tmp/go1.20.5.darwin-arm64.tar.gz -sSL https://go.dev/dl/go1.20.5.darwin-arm64.tar.gz
echo "94ad76b7e1593bb59df7fd35a738194643d6eed26a4181c94e3ee91381e40459 /tmp/go1.20.5.darwin-arm64.tar.gz" | sha256sum -c -
mkdir /usr/local/goenv/1.20.5
tar --strip-components=1 -xzf /tmp/go1.20.5.darwin-arm64.tar.gz -C /usr/local/goenv/1.20.5
ln -s /usr/local/goenv/1.20.5 /usr/local/go`,
		},
		{
			name:       "beepboop",
			rootDir:    "/usr/local/go",
			installDir: "usr/local/go",
			version:    semver.MustParse("1.20.5"),
			want:       `curl -o /tmp/go1.20.5.darwin-arm64.tar.gz -sSL https://go.dev/dl/go1.20.5.darwin-arm64.tar.gz
echo "94ad76b7e1593bb59df7fd35a738194643d6eed26a4181c94e3ee91381e40459 /tmp/go1.20.5.darwin-arm64.tar.gz" | sha256sum -c -
mkdir usr/local/go
tar --strip-components=1 -xzf /tmp/go1.20.5.darwin-arm64.tar.gz -C usr/local/go`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := InstallScript(tt.rootDir, tt.installDir, tt.version)
			assert.Equal(t, tt.want, actual)
		})
	}
}
