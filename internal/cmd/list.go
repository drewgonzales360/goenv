package cmd

import (
	"fmt"
	"os"

	"github.com/Masterminds/semver"
	"github.com/urfave/cli/v2"
)

type versionRange struct {
	minVerson semver.Version
	maxVerson semver.Version
}

var goVersion map[string][]string = map[string][]string{
	"1.18": {
		"1.18.1",
		"1.18",
	},
	"1.17": {
		"1.17.9",
		"1.17.8",
		"1.17.7",
		"1.17.6",
		"1.17.5",
		"1.17.4",
		"1.17.3",
		"1.17.2",
		"1.17.1",
		"1.17",
	},
	"1.16": {
		"1.16.15",
		"1.16.14",
		"1.16.13",
		"1.16.12",
		"1.16.11",
		"1.16.10",
		"1.16.9",
		"1.16.8",
		"1.16.7",
		"1.16.6",
		"1.16.5",
		"1.16.4",
		"1.16.3",
		"1.16.2",
		"1.16.1",
		"1.16",
	},
	"1.15": {
		"1.15.15",
		"1.15.14",
		"1.15.13",
		"1.15.12",
		"1.15.11",
		"1.15.10",
		"1.15.9",
		"1.15.8",
		"1.15.7",
		"1.15.6",
		"1.15.5",
		"1.15.4",
		"1.15.3",
		"1.15.2",
		"1.15.1",
		"1.15",
	},
}

func ListCommand(c *cli.Context) error {
	versions, err := os.ReadDir(InstallDirectory)
	if err != nil {
		return err
	}

	if c.Bool("available") {
		fmt.Println("Installed Versions:")
	}
	for _, version := range versions {
		fmt.Println(version.Name())
	}

	if c.Bool("available") {
		fmt.Println("Available Versions:")
		for k, v := range goVersion {
			fmt.Printf("%s: %v\n", k, v)
		}
	}

	return nil
}
