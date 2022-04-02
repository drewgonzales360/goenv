package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/drewgonzales360/goenv/pkg"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

const (
	UsrLocalBin string = "/usr/local/bin/"
	PermError   string = "are you sure you're root? you do not have access to %v"
)

func UseCommand(c *cli.Context) error {
	version := ""
	if c.NArg() > 0 {
		version = c.Args().Get(0)
	}

	if err := Use(version); err != nil {
		pkg.Fail(err.Error())
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

	if err = Link(goVersion, "go"); err != nil {
		return err
	}

	if err = Link(goVersion, "gofmt"); err != nil {
		return err
	}

	output, err := exec.Command("/usr/local/bin/go", "version").Output()
	if err != nil {
		pkg.Debug(err.Error())
		return err
	}

	pkg.Success(fmt.Sprintf("Using %s", strings.TrimSuffix(string(output), "\n")))
	return nil
}

func Link(goVersion *semver.Version, binary string) error {
	usrLocalBinSymlink := UsrLocalBin + binary
	if _, err := os.Stat(usrLocalBinSymlink); err == nil {
		if err = os.Remove(usrLocalBinSymlink); err != nil {
			return errors.Wrap(err, "could not remove "+usrLocalBinSymlink)
		}
	}

	usrLocalGoVersionBin := InstallDirectory + goVersion.Original() + "/bin/" + binary
	if _, err := os.Stat(usrLocalGoVersionBin); err != nil {
		pkg.Debug(err.Error())
		return fmt.Errorf("could not find go version %s. goenv install %s", goVersion.Original(), goVersion.Original())
	}

	if err := os.Symlink(usrLocalGoVersionBin, usrLocalBinSymlink); err != nil {
		return errors.Wrap(err, "could not link "+binary)
	}

	return nil
}
