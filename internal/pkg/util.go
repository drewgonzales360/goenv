package pkg

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"runtime"
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

	os := runtime.GOOS
	arch := runtime.GOARCH
	if os != "linux" && os != "darwin" {
		Debug(fmt.Sprintf("Running an unsupported os: %s", os))
	}
	if arch != "amd64" {
		Debug(fmt.Sprintf("Running an unsupported arch: %s", arch))
	}

	url := fmt.Sprintf("https://go.dev/dl/go%s.%s-%s.tar.gz", urlVersion, os, arch)
	Debug(fmt.Sprintf("Downloading %s", url))
	return url
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

	if ok, err := checkHash(filepath, getHash(v)); err != nil || !ok {
		return "", err
	}

	s.FinalMSG = fmt.Sprintf("✅ Downloaded Go %s\n", v.Original())
	return filepath, nil
}

// checkHash will check if the hash of the file matches the hash advertised on
// go.dev/dl. If we have the hash written down in our code, we'll check it against
// what we downloaded. If we haven't put the hash in, this won't throw an error.
// This way, we don't _have_ to update every time a new version comes out. We
// just won't check the hash.
func checkHash(file string, expected string) (bool, error) {
	if expected == "" {
		return true, nil
	}
	out, err := os.ReadFile(file)
	if err != nil {
		return false, err
	}

	sum := sha256.Sum256(out)
	downloaded := hex.EncodeToString(sum[:])
	if downloaded != expected {
		return false, fmt.Errorf("file corrupted, downloaded: %s, expected %s", downloaded, expected)
	}

	return true, nil
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
