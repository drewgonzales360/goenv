package cmd

import (
	"fmt"
	"os"
	"os/exec"
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

	if err := Use(version); err != nil {
		return err
	}
	return nil
}

func Use(version string) error {
	if inaccessible := pkg.CheckRW(); len(inaccessible) > 0 {
		return fmt.Errorf(PermError, inaccessible)
	}

	goVersion, err := semver.NewVersion(version)
	if err != nil {
		return errors.Wrap(err, "could not parse version as a semver")
	}

	if err = link(goVersion); err != nil {
		return err
	}

	output, err := exec.Command(UsrLocalGo+"/bin/go", "version").Output()
	if err != nil {
		pkg.Debug(err.Error())
		return err
	}

	pkg.Success(fmt.Sprintf("Using %s", strings.TrimSuffix(string(output), "\n")))
	return nil
}

func link(goVersion *semver.Version) error {
	if _, err := os.Stat(UsrLocalGo); err == nil {
		if err = os.Remove(UsrLocalGo); err != nil {
			return errors.Wrap(err, "could not remove "+UsrLocalGo)
		}
	}

	usrLocalGoVersion := InstallDirectory + goVersion.Original()
	if _, err := os.Stat(usrLocalGoVersion); err != nil {
		pkg.Debug(err.Error())
		return fmt.Errorf("could not find go version %s. goenv install %s", goVersion.Original(), goVersion.Original())
	}

	if err := os.Symlink(usrLocalGoVersion, UsrLocalGo); err != nil {
		return errors.Wrap(err, "could not link")
	}

	return nil
}
