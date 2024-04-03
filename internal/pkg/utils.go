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
	"os"
	"path"
	"time"
)

// CheckRW checks if the user can install and use Go by making sure they have read and write access
// to the directories where go will be installed.
func CheckRW(goenvRootDirectory string, goenvInstallDirectory string) []string {
	accessDenied := []string{}
	currentTime := time.Now().Local()

	if err := removeDeadLink(goenvRootDirectory); err != nil {
		Debug(err.Error())
	}

	dirs := []string{goenvRootDirectory, goenvInstallDirectory}
	for _, dir := range dirs {
		if _, err := os.Stat(dir); err == nil {
			if err := os.Chtimes(dir, currentTime, currentTime); err != nil {
				Debug(err.Error())
				accessDenied = append(accessDenied, dir)
			}
		} else {
			if err := os.MkdirAll(dir, 0755); err != nil {
				Debug(err.Error())
				accessDenied = append(accessDenied, dir)
			}
		}
	}

	return accessDenied
}

// removeDeadLink checks if the go root is pointed at an uninstalled go version. If that's the case,
// then we remove the link.
func removeDeadLink(path string) error {
	installDir, err := os.Readlink(path)
	if err != nil {
		return err
	}

	_, err = os.Stat(installDir)
	if err == nil {
		return nil
	}

	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

// CheckInstalled simply checks if a directory exists.
func CheckInstalled(goenvInstallDirectory string, version string) error {
	_, err := os.Stat(path.Join(goenvInstallDirectory, version))
	return err
}
