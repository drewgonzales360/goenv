package pkg

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/fatih/color"
)

type GoVersion struct {
	DarwinHash string
	LinuxHash  string
}

type GoVersionList = map[string]map[string]GoVersion

// CreateGoVersionList will take a list of all the directories in
// goenv and then parse it into a GoVersionList. We do this so that
// we can print the installed Go versions nicely for the user.
func CreateGoVersionList(directories []string) GoVersionList {
	goVersionList := GoVersionList{}
	for _, d := range directories {
		s := semver.MustParse(d)
		majorMinor := fmt.Sprintf("%d.%d", s.Major(), s.Minor())
		if _, ok := goVersionList[majorMinor]; !ok {
			goVersionList[majorMinor] = make(map[string]GoVersion)
		}
		goVersionList[majorMinor][s.Original()] = GoVersion{}
	}

	return goVersionList
}

// Print will nicely display the installed and available versions of
// Go you can install to the terminal.
func Print(g *GoVersionList) {
	for minorVersion, patchVersions := range *g {
		color.New(color.FgHiBlack).Printf("%s: ", minorVersion)

		patches := []string{}
		for key := range patchVersions {
			patches = append(patches, key)
		}
		fmt.Println(strings.Join(patches, " "))
	}
}

// GetHash reads through our database in hashes.go and returns the
// known shasum. If we don't know the shasum, either because the
// version of Go is too new or too old, we return an empty string.
func GetHash(v *semver.Version) string {
	majorMinor := fmt.Sprintf("%d.%d", v.Major(), v.Minor())

	d, ok := GoVersions[majorMinor][v.Original()]
	if !ok {
		Debug("unknown version, but let's download it anyway")
		return ""
	}

	if runtime.GOOS == "darwin" {
		return d.DarwinHash
	} else if runtime.GOOS == "linux" {
		return d.LinuxHash
	}

	Debug("unknown os, but let's download it anyway")
	return ""
}
