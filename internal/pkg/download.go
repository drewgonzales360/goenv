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
	"strings"
	"time"

	"github.com/Masterminds/semver"
	"github.com/briandowns/spinner"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	tempDir = "/tmp/goenv"
)

// randomString returns 6 random characters from charset.
func randomString() string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// DownloadFile will download a url to a local file. If we have a recorded shasum for the
// file, we'll check it. We'll return a path to the downloaded file.
// https://go.dev/dl/go1.18.linux-amd64.tar.gz
func DownloadFile(v *semver.Version) (filepath string, err error) {
	s := spinner.New(spinner.CharSets[38], 200*time.Millisecond)
	s.Suffix = fmt.Sprintf(" Downloading Go %s", v.Original()) // Build our new spinner
	s.Start()                                                  // Start the spinner
	defer s.Stop()

	url, checksum := getDownloadInfo(v)
	client := http.Client{Timeout: time.Second * HttpRequestTimeout}
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("could not download file from %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("could not dowload go - response code: %s", resp.Status)
	}

	// Create the file
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		return "", err
	}

	filepath = fmt.Sprintf("%s/%s.tar.gz", tempDir, randomString())
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

	validSha, err := checkHash(filepath, checksum)
	if err != nil || !validSha {
		return "", err
	}

	s.FinalMSG = fmt.Sprintf("✅ Downloaded and validated Go %s\n", v.Original())
	if !validSha {
		s.FinalMSG = fmt.Sprintf("✅ Downloaded Go %s\n", v.Original())
	}

	return filepath, nil
}

// checkHash will check if the hash of the file matches the hash advertised on
// go.dev/dl. If we have the hash written down in our code, we'll check it against
// what we downloaded. If we haven't put the hash in, this won't throw an error.
// This way, we don't _have_ to update every time a new version comes out. We
// just won't check the hash.
func checkHash(file string, expected ChecksumSHA256) (bool, error) {
	if expected == "" {
		return true, nil
	}
	out, err := os.ReadFile(file)
	if err != nil {
		return false, err
	}

	sum := sha256.Sum256(out)
	downloaded := hex.EncodeToString(sum[:])
	if downloaded != string(expected) {
		return false, fmt.Errorf("file corrupted, downloaded: %s, expected %s", downloaded, expected)
	}

	return true, nil
}

// ExtractTarGz will do the equivalent of a tar -xzvf -C and untar the
// tarball to whichever destination path we need to go to.
func ExtractTarGz(tarballPath, destinationPath string) error {
	s := spinner.New(spinner.CharSets[38], 200*time.Millisecond)
	s.Suffix = " Extracting package" // Build our new spinner
	s.Start()                        // Start the spinner
	defer s.Stop()

	r, err := os.Open(tarballPath)
	if err != nil {
		return fmt.Errorf("could not open tarball: %w", err)
	}

	uncompressedStream, err := gzip.NewReader(r)
	if err != nil {
		return fmt.Errorf("could not unzip file: %w", err)
	}
	defer uncompressedStream.Close()

	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("could not extract file: %w", err)
		}

		extractionDestination := fmt.Sprintf("%s/%s", destinationPath, strings.TrimPrefix(header.Name, "go/"))
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(extractionDestination, 0755); err != nil {
				return fmt.Errorf("could not create directory: %w", err)
			}
		case tar.TypeReg:
			outFile, err := os.Create(extractionDestination)
			if err != nil {
				return fmt.Errorf("could not create file: %w", err)
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return fmt.Errorf("could not write to file: %w", err)
			}

			err = os.Chmod(outFile.Name(), header.FileInfo().Mode())
			if err != nil {
				return fmt.Errorf("could not set extracted file permissions: %w", err)
			}

			outFile.Close()

		default:
			return fmt.Errorf(
				"unknown type: %s in %s",
				header.Format,
				header.Name)
		}
	}

	s.FinalMSG = "✅ Extracted package\n"
	return nil
}
