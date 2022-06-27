package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/drewgonzales360/goenv/internal/pkg"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func UseCommand(c *cli.Context) error {
	version, err := parseVersionArg(c)
	if err != nil {
		return err
	}

	config, err := parseConfig(c)
	if err != nil {
		return err
	}

	if err := Use(config, version); err != nil {
		return err
	}
	return nil
}

func Use(config *pkg.Config, version string) error {
	if inaccessible := pkg.CheckRW(config); len(inaccessible) > 0 {
		return fmt.Errorf(PermError, inaccessible)
	}

	goVersion, err := semver.NewVersion(version)
	if err != nil {
		return errors.Wrap(err, "could not parse version as a semver")
	}

	if err = link(config, goVersion); err != nil {
		return err
	}

	output, err := exec.Command("go", "version").Output()
	if err != nil {
		pkg.Debug(err.Error())
		return err
	}

	pkg.Success(fmt.Sprintf("Using %s", strings.TrimSuffix(string(output), "\n")))
	return nil
}

func link(config *pkg.Config, goVersion *semver.Version) error {
	if _, err := os.Stat(config.GoenvInstallDirectory); err == nil {
		if err = os.Remove(config.GoenvInstallDirectory); err != nil {
			return errors.Wrap(err, "could not remove "+config.GoenvInstallDirectory)
		}
	}

	usrLocalGoVersion := path.Join(config.GoenvRootDirectory, goVersion.Original())
	if _, err := os.Stat(usrLocalGoVersion); err != nil {
		pkg.Debug(err.Error())
		return fmt.Errorf("could not find go version %s. goenv install %s", goVersion.Original(), goVersion.Original())
	}

	if err := os.Symlink(usrLocalGoVersion, config.GoenvInstallDirectory); err != nil {
		return errors.Wrap(err, "could not link")
	}

	return nil
}
