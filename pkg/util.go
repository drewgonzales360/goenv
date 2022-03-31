package pkg

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/Masterminds/semver"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	tempDir = "/tmp/goenv"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString() string {
	return StringWithCharset(8, charset)
}

func FormatDownloadURL(v semver.Version) string {
	return fmt.Sprintf("https://go.dev/dl/go%s.linux-amd64.tar.gz", v.String())
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
// https://go.dev/dl/go1.18.linux-amd64.tar.gz
func DownloadFile(url string) (string, error) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Create the file
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		return "", err
	}

	filepath := fmt.Sprintf("/tmp/goenv/%s.tar.gz", RandomString())
	out, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return filepath, err
}

func ExtractTarGz(tarballPath, destinationPath string) error {
	r, err := os.Open(tarballPath)
	if err != nil {
		return errors.Wrap(err, "could not open tarball")
	}

	uncompressedStream, err := gzip.NewReader(r)
	if err != nil {
		return errors.Wrap(err, "could not unzip file")
	}
	defer uncompressedStream.Close()

	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrap(err, "could not extract file")
		}

		extractionDestination := fmt.Sprintf("%s/%s", destinationPath, strings.TrimPrefix(header.Name, "go/"))
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(extractionDestination, 0755); err != nil {
				return errors.Wrap(err, "could not create directory")
			}
		case tar.TypeReg:
			outFile, err := os.Create(extractionDestination)
			if err != nil {
				return errors.Wrap(err, "could not create file")
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return errors.Wrap(err, "could not write to file")
			}
			outFile.Close()

		default:
			return errors.New(fmt.Sprintf(
				"ExtractTarGz: uknown type: %s in %s",
				header.Format,
				header.Name))
		}
	}
	return nil
}
