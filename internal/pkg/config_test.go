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
	"os"
	"reflect"
	"testing"

	"github.com/drewgonzales360/goenv/internal/pkg"
)

func TestConfigReading(t *testing.T) {
	expectedRootDir := "/home/gopher/.local/go"
	expectedInstallDir := "/home/gopher/.local/goenv"
	if err := os.Setenv(pkg.GoEnvRootDirEnvVar, expectedRootDir); err != nil {
		t.Fail()
	}
	if err := os.Setenv(pkg.GoEnvInstallDirEnvVar, expectedInstallDir); err != nil {
		t.Fail()
	}

	leftConfig := pkg.ReadConfig()
	rightConfig := &pkg.Config{
		expectedRootDir,
		expectedInstallDir,
	}
	if !reflect.DeepEqual(leftConfig, rightConfig) {
		t.Fail()
	}
}
