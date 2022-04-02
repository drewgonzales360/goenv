package pkg

import (
	"testing"

	"github.com/Masterminds/semver"
)

func TestEverything(t *testing.T) {

	tarballPath, err := DownloadFile(*semver.MustParse("1.19"))
	if err != nil {
		t.Fatal(err)
	}

	if err = ExtractTarGz(tarballPath, "/tmp/goenv/go1.18"); err != nil {
		t.Fatal(err)
	}
}
