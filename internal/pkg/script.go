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
	"path"

	"github.com/Masterminds/semver/v3"
)

func InstallScript(rootDir string, installDir string, version *semver.Version) string {
	url, checksum := getDownloadInfo(version)
	localTarball := path.Join("/tmp", path.Base(url))

	installScriptTemplate := `curl -o %s -sSL %s
echo "%s %s" | sha256sum -c -
mkdir -p %s
tar --strip-components=1 -xzf %s -C %s`

	script := fmt.Sprintf(
		installScriptTemplate,
		localTarball,
		url,
		checksum,
		localTarball,
		installDir,
		localTarball,
		installDir,
	)

	if rootDir != installDir {
		script += fmt.Sprintf("\nln -s %s %s", installDir, rootDir)
	}

	script += fmt.Sprintf("\nexport PATH=%s/bin:$PATH", rootDir)

	return script
}
