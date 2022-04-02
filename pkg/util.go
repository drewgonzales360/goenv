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

	"github.com/Masterminds/semver"
	"github.com/briandowns/spinner"
	"github.com/pkg/errors"
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
	urlVersion := v.String()
	// If we have 1.18, we'd parse the version to 1.18.0, but the URL doesn't
	// actually inclued the last .0
	if v.Patch() == 0 {
		urlVersion = strings.TrimSuffix(urlVersion, ".0")
	}
	return fmt.Sprintf("https://go.dev/dl/go%s.linux-amd64.tar.gz", urlVersion)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
// https://go.dev/dl/go1.18.linux-amd64.tar.gz
func DownloadFile(v semver.Version) (filepath string, err error) {
	s := spinner.New(spinner.CharSets[38], 200*time.Millisecond)
	s.Suffix = fmt.Sprintf(" Downloading Go %s", v.Original()) // Build our new spinner
	s.Start()                                                  // Start the spinner
	defer s.Stop()

	url := FormatDownloadURL(v)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("could not dowload go - response code: %s", resp.Status)
	}

	// Create the file
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		return "", err
	}

	filepath = fmt.Sprintf("/tmp/goenv/%s.tar.gz", RandomString())
	out, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	s.FinalMSG = fmt.Sprintf("✅ Downloaded Go %s\n", v.Original())
	return filepath, nil
}

func ExtractTarGz(tarballPath, destinationPath string) error {
	s := spinner.New(spinner.CharSets[38], 200*time.Millisecond)
	s.Suffix = " Extracting package" // Build our new spinner
	s.Start()                        // Start the spinner
	defer s.Stop()

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

			err = os.Chmod(outFile.Name(), header.FileInfo().Mode())
			if err != nil {
				return errors.Wrap(err, "could not set extracted file permissions")
			}

			outFile.Close()

		default:
			return errors.New(fmt.Sprintf(
				"ExtractTarGz: uknown type: %s in %s",
				header.Format,
				header.Name))
		}
	}

	s.FinalMSG = "✅ Extracted package\n"
	return nil
}
