package pkg

import (
	"testing"

	"github.com/Masterminds/semver"
)

func TestDownloadAndUntar(t *testing.T) {
	version := "1.18"
	goVersion := semver.MustParse(version)
	tarballPath, err := DownloadFile(*goVersion)
	if err != nil {
		t.Fatal(err)
	}

	if err = ExtractTarGz(tarballPath, "/tmp/goenv/"+version); err != nil {
		t.Fatal(err)
	}
}
