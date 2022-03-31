package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Masterminds/semver"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func Use(c *cli.Context) error {
	version := ""
	if c.NArg() > 0 {
		version = c.Args().Get(0)
	}

	goVersion, err := semver.NewVersion(version)
	if err != nil {
		return fmt.Errorf("could not parse version as a semver")
	}

	if err := os.Symlink(InstallDirectory+goVersion.Original()+"/bin/go", "/usr/local/bin/go"); err != nil {
		return errors.Wrap(err, "could not link go")
	}

	if err := os.Symlink(InstallDirectory+goVersion.Original()+"/bin/gofmt", "/usr/local/bin/gofmt"); err != nil {
		return errors.Wrap(err, "could not link gofmt")
	}

	output, _ := exec.Command("go version").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(output))

	return nil
}
