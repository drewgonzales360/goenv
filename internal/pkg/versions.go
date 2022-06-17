package pkg

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/fatih/color"
)

type GoVersion struct {
	Version string
	hash    string
}

type GoVersionList = map[string][]GoVersion

func CreateGoVersionList(directories []string) GoVersionList {
	goVersionList := GoVersionList{}
	for _, d := range directories {
		s := semver.MustParse(d)
		majorMinor := fmt.Sprintf("%d.%d", s.Major(), s.Minor())
		if _, ok := goVersionList[majorMinor]; ok {
			goVersionList[majorMinor] = append(goVersionList[majorMinor], GoVersion{d, ""})
		} else {
			goVersionList[majorMinor] = []GoVersion{{d, ""}}
		}
	}

	return goVersionList
}

func Print(g *GoVersionList) {
	for minorVersion, patchVersions := range *g {
		color.New(color.FgHiBlack).Printf("%s: ", minorVersion)

		patches := []string{}
		for _, p := range patchVersions {
			patches = append(patches, p.Version)
		}
		fmt.Println(strings.Join(patches, " "))
	}
}

var GoVersions GoVersionList = GoVersionList{
	"1.18": {
		{"1.18.3", "adf"},
		{"1.18.2", "adf"},
		{"1.18.1", "adf"},
		{"1.18", "adf"},
	},
	"1.17": {
		{"1.17.12", "adf"},
		{"1.17.11", "adf"},
		{"1.17.10", "adf"},
		{"1.17.9", "adf"},
		{"1.17.8", "adf"},
		{"1.17.7", "adf"},
		{"1.17.6", "adf"},
		{"1.17.5", "adf"},
		{"1.17.4", "adf"},
		{"1.17.3", "adf"},
		{"1.17.2", "adf"},
		{"1.17.1", "adf"},
		{"1.17", "adf"},
	},
	"1.16": {
		{"1.16.12", "adf"},
		{"1.16.11", "adf"},
		{"1.16.10", "adf"},
		{"1.16.9", "adf"},
		{"1.16.8", "adf"},
		{"1.16.7", "adf"},
		{"1.16.6", "adf"},
		{"1.16.5", "adf"},
		{"1.16.4", "adf"},
		{"1.16.3", "adf"},
		{"1.16.2", "adf"},
		{"1.16.1", "adf"},
		{"1.16", "adf"},
	},
}
