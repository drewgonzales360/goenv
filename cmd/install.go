package cmd

import (
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/drewgonzales360/goenv/pkg"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

const (
	InstallDirectory string = "/usr/local/go/"
)

func Install(c *cli.Context) error {
	version := ""
	if c.NArg() > 0 {
		version = c.Args().Get(0)
	}

	goVersion, err := semver.NewVersion(version)
	if err != nil {
		return fmt.Errorf("could not parse version as a semver")
	}

	downloadURL := pkg.FormatDownloadURL(*goVersion)
	filePath, err := pkg.DownloadFile(downloadURL)
	if err != nil {
		return errors.Wrap(err, "could not download go")
	}

	err = pkg.ExtractTarGz(filePath, InstallDirectory+goVersion.Original())
	if err != nil {
		return errors.Wrap(err, "could not extract go")
	}

	Use(c)
	return nil
}
