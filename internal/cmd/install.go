package cmd

import (
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/drewgonzales360/goenv/internal/pkg"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func InstallCommand(c *cli.Context) error {
	version, err := parseVersionArg(c)
	if err != nil {
		return err
	}

	if err := Install(version); err != nil {
		return err
	}
	return nil
}

func Install(version string) error {
	if inaccessible := pkg.CheckRW(); len(inaccessible) > 0 {
		return fmt.Errorf(PermError, inaccessible)
	}

	goVersion, err := semver.NewVersion(version)
	if err != nil {
		return fmt.Errorf("could not parse version as a semver")
	}

	filePath, err := pkg.DownloadFile(*goVersion)
	if err != nil {
		return errors.Wrap(err, "could not download go")
	}

	err = pkg.ExtractTarGz(filePath, InstallDirectory+goVersion.Original())
	if err != nil {
		return errors.Wrap(err, "could not extract go")
	}

	return Use(version)
}
