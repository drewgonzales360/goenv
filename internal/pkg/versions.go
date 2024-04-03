package pkg

// ///////////////////////////////////////////////////////////////////////
// Copyright 2024 Drew Gonzales
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
import (
	"fmt"
	"sort"
	"strings"

	"github.com/Masterminds/semver/v3"
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

// Print will nicely display the installed and available versions of Go you can install to the
// terminal.
func Print(g map[string][]string, linePrefix string) {
	keys := make([]string, 0, len(g))
	for k := range g {
		keys = append(keys, k)
	}
	m := sortSemvers(keys)

	for _, key := range m {
		color.New(color.FgHiBlack).Printf("%s%s: ", linePrefix, key.Original())

		l := sortSemvers(g[key.Original()])
		sa := mapToOriginal(l)
		fmt.Println(strings.Join(sa, " "))
	}
}

// sortSemvers maps the raw strings into semver.Versions then sorts it.
func sortSemvers(raw []string) []*semver.Version {
	vs := make([]*semver.Version, len(raw))
	for i, r := range raw {
		v, err := semver.NewVersion(r)
		if err != nil {
			Error(fmt.Sprintf("could not parse semver: %s %s", err.Error(), r))
			continue
		}

		vs[i] = v
	}

	sort.Sort(semver.Collection(vs))
	return vs
}

// mapToOriginal turns the semver.Versions back into strings.
func mapToOriginal(vs []*semver.Version) (sa []string) {
	for _, v := range vs {
		sa = append(sa, v.Original())
	}
	return sa
}
