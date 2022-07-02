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
