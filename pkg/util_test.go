package util

import "testing"

func TestEverything(t *testing.T) {
	tarballPath, err := DownloadFile("https://go.dev/dl/go1.18.linux-amd64.tar.gz")
	if err != nil {
		t.Fatal(err)
	}

	if err = ExtractTarGz(tarballPath, "/tmp/goenv/go1.18"); err != nil {
		t.Fatal(err)
	}
}
