package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/Masterminds/semver"
	"github.com/drewgonzales360/goenv/internal/pkg"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func UninstallCommand(c *cli.Context) error {
	version, err := parseVersionArg(c)
	if err != nil {
		return err
	}

	config, err := parseConfig(c)
	if err != nil {
		return err
	}

	if err := Uninstall(config, version); err != nil {
		return err
	}
	return nil
}

func Uninstall(config *pkg.Config, version string) error {
	if inaccessible := pkg.CheckRW(config); len(inaccessible) > 0 {
		return fmt.Errorf(PermError, inaccessible)
	}

	goVersion, err := semver.NewVersion(version)
	if err != nil {
		return errors.Wrap(err, "could not parse version as a semver")
	}

	if err = os.RemoveAll(path.Join(config.GoenvRootDirectory, goVersion.Original())); err != nil {
		return errors.Wrap(err, "could not uninstall go")
	}

	pkg.Success("Uninstalled Go " + goVersion.Original())
	return nil
}
