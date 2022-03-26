package cmd

import (
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/urfave/cli/v2"
)

const (
	InstallDirectory string = "/usr/local/go"
)

func Install(c *cli.Context) error {
	version := ""
	if c.NArg() > 0 {
		version = c.Args().Get(0)
	}

	_, err := semver.NewVersion(version)
	if err != nil {
		return fmt.Errorf("could not parse version")
	}

	// Download go tarball

	// Untar

	return nil
}
