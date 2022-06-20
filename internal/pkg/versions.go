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

func getHash(v semver.Version) string {
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

var GoVersions GoVersionList = GoVersionList{
	"1.18": {
		"1.18.3": GoVersion{
			DarwinHash: "d9dcf8fc35da54c6f259be41954783a9f4984945a855d03a003a7fd6ea4c5ca1",
		},
		"1.18.2": GoVersion{},
		"1.18.1": GoVersion{},
		"1.18":   GoVersion{},
	},
	"1.17": {
		"1.17.12": GoVersion{},
		"1.17.11": GoVersion{},
		"1.17.10": GoVersion{},
		"1.17.9":  GoVersion{},
		"1.17.8":  GoVersion{},
		"1.17.7":  GoVersion{},
		"1.17.6":  GoVersion{},
		"1.17.5":  GoVersion{},
		"1.17.4":  GoVersion{},
		"1.17.3":  GoVersion{},
		"1.17.2":  GoVersion{},
		"1.17.1":  GoVersion{},
		"1.17":    GoVersion{},
	},
	"1.16": {
		"1.16.12": GoVersion{},
		"1.16.11": GoVersion{},
		"1.16.10": GoVersion{},
		"1.16.9":  GoVersion{},
		"1.16.8":  GoVersion{},
		"1.16.7":  GoVersion{},
		"1.16.6":  GoVersion{},
		"1.16.5":  GoVersion{},
		"1.16.4":  GoVersion{},
		"1.16.3":  GoVersion{},
		"1.16.2":  GoVersion{},
		"1.16.1":  GoVersion{},
		"1.16":    GoVersion{},
	},
}
