package cmd

import (
	"fmt"
	"os"

	"github.com/Masterminds/semver"
	"github.com/drewgonzales360/goenv/pkg"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func UninstallCommand(c *cli.Context) error {
	version, err := parseVersionArg(c)
	if err != nil {
		return err
	}

	if err := Uninstall(version); err != nil {
		return err
	}
	return nil
}

func Uninstall(version string) error {
	if inaccessible := pkg.CheckRW(); len(inaccessible) > 0 {
		return fmt.Errorf(PermError, inaccessible)
	}

	goVersion, err := semver.NewVersion(version)
	if err != nil {
		return errors.Wrap(err, "could not parse version as a semver")
	}

	if err = os.RemoveAll(InstallDirectory + goVersion.Original()); err != nil {
		return errors.Wrap(err, "could not uninstall go")
	}

	pkg.Success("Uninstalled Go " + goVersion.Original())
	return nil
}
