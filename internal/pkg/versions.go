package pkg

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/fatih/color"
)

// CreateGoVersionList will take a list of all the directories in
// goenv and then parse it into a GoVersionList. We do this so that
// we can print the installed Go versions nicely for the user.
func CreateGoVersionList(directories []string) map[string][]string {
	goVersionList := make(map[string][]string)
	for _, d := range directories {
		s, err := semver.NewVersion(d)
		if err != nil {
			Debug(fmt.Sprintf("could not parse semver: %s %s", err.Error(), d))
			continue
		}

		majorMinor := fmt.Sprintf("%d.%d", s.Major(), s.Minor())
		if _, ok := goVersionList[majorMinor]; !ok {
			goVersionList[majorMinor] = []string{}
		}
		goVersionList[majorMinor] = append(goVersionList[majorMinor], s.Original())
	}

	return goVersionList
}

// Print will nicely display the installed and available versions of
// Go you can install to the terminal.
func Print(g map[string][]string) {
	keys := make([]string, 0, len(g))
	for k := range g {
		keys = append(keys, k)
	}
	m := sortSemvers(keys)

	for _, key := range m {
		color.New(color.FgHiBlack).Printf("%s: ", key.Original())

		l := sortSemvers(g[key.Original()])
		sa := mapToOriginal(l)
		fmt.Println(strings.Join(sa, " "))
	}
}

func sortSemvers(raw []string) []*semver.Version {
	vs := make([]*semver.Version, len(raw))
	for i, r := range raw {
		v, err := semver.NewVersion(r)
		if err != nil {
			Fail(fmt.Sprintf("could not parse semver: %s %s", err.Error(), r))
		}

		vs[i] = v
	}

	sort.Sort(semver.Collection(vs))
	return vs
}

func mapToOriginal(vs []*semver.Version) (sa []string) {
	for _, v := range vs {
		sa = append(sa, v.Original())
	}
	return sa
}
