package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/Masterminds/semver"
	"github.com/pkg/errors"
)

type ChecksumSHA256 string

// This comes from https://github.com/golang/website/blob/master/internal/dl/dl.go
type Release struct {
	Version        string `json:"version"`
	Stable         bool   `json:"stable"`
	Files          []File `json:"files"`
	Visible        bool   `json:"-"` // show files on page load
	SplitPortTable bool   `json:"-"` // whether files should be split by primary/other ports.
}

// This comes from https://github.com/golang/website/blob/master/internal/dl/dl.go
type File struct {
	Filename       string    `json:"filename"`
	OS             string    `json:"os"`
	Arch           string    `json:"arch"`
	Version        string    `json:"version"`
	Checksum       string    `json:"-" datastore:",noindex"` // SHA1; deprecated
	ChecksumSHA256 string    `json:"sha256" datastore:",noindex"`
	Size           int64     `json:"size" datastore:",noindex"`
	Kind           string    `json:"kind"` // "archive", "installer", "source"
	Uploaded       time.Time `json:"-"`
}

// GetGoVersions queries the go.dev/dl for the current releases. The latest patch
// versions of the latest two minor versions are considered stable. If you want
// all releases of Go, pass in true.
func GetGoVersions(getAllVersions bool) ([]Release, error) {
	includeAll := "&include=all"
	goDevURL := "https://go.dev/dl/?mode=json"
	if getAllVersions {
		goDevURL = goDevURL + includeAll
	}

	resp, err := http.DefaultClient.Get(goDevURL)
	if err != nil {
		return nil, errors.Wrap(err, "could not get go versions")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("go.dev returned %s", resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var releases []Release
	err = json.Unmarshal(bodyBytes, &releases)
	if err != nil {
		return nil, err
	}

	return releases, nil
}

// getDownloadInfo will make a best effort to query go.dev for the latest releases
// and their shasums. If we can't reach it, default to the well known formats
// and locations of the tarballs.
func getDownloadInfo(v *semver.Version) (url string, checkSum ChecksumSHA256) {
	releases, err := GetGoVersions(true)
	if err != nil {
		Debug("could not query go.dev for shasums, defaulting to best effort")
		return defaultDownloadURL(v), ""
	}

	urlVersion := fmt.Sprintf("go%s", toLooseGoVersion(v))
	for _, release := range releases {
		if release.Version == urlVersion {
			for _, file := range release.Files {
				if file.Arch == runtime.GOARCH && file.OS == runtime.GOOS && file.Kind == "archive" {
					checkSum = ChecksumSHA256(file.ChecksumSHA256)
					url = fmt.Sprintf("https://go.dev/dl/%s", file.Filename)
				}
			}
		}
	}

	if url == "" {
		Error("could not find release")
	}
	return url, checkSum
}

// defaultDownloadURL assumes the format and URL of the tarball we
// need to install Go.
func defaultDownloadURL(v *semver.Version) string {
	urlVersion := toLooseGoVersion(v)

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

// toLooseGoVersion will take a strict semver and turn it into the looser Go convention of
// stripping the minor patch when it's .0.
func toLooseGoVersion(v *semver.Version) string {
	urlVersion := v.String()
	// If we have 1.18, we'd parse the version to 1.18.0, but the URL doesn't
	// actually inclued the last .0
	if v.Patch() == 0 {
		urlVersion = strings.TrimSuffix(urlVersion, ".0")
	}

	return urlVersion
}

func ListAvailableVersions(getAllVersions bool) ([]string, error) {
	releases, err := GetGoVersions(getAllVersions)
	if err != nil {
		return nil, err
	}

	versions := []string{}
	for _, release := range releases {
		version := strings.TrimPrefix(release.Version, "go")
		versions = append(versions, version)
	}

	return versions, nil
}
