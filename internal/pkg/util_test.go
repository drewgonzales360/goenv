package pkg_test

import (
	"testing"

	"github.com/Masterminds/semver"
	"github.com/drewgonzales360/goenv/internal/pkg"
)

func TestDownloadAndUntar(t *testing.T) {
	version := "1.18"
	goVersion := semver.MustParse(version)
	tarballPath, err := pkg.DownloadFile(*goVersion)
	if err != nil {
		t.Fatal(err)
	}

	if err = pkg.ExtractTarGz(tarballPath, "/tmp/goenv/"+version); err != nil {
		t.Fatal(err)
	}
}
