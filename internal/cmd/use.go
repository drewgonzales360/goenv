package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/urfave/cli/v2"

	"github.com/drewgonzales360/goenv/internal/pkg"
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

	if err := use(config, version); err != nil {
		return err
	}
	return nil
}

func use(config *pkg.Config, version string) error {
	if inaccessible := pkg.CheckRW(config); len(inaccessible) > 0 {
		return fmt.Errorf(PermError, inaccessible)
	}

	goVersion, err := semver.NewVersion(version)
	if err != nil {
		return fmt.Errorf("could not parse version as a semver: %w", err)
	}

	if err = link(config, goVersion); err != nil {
		return err
	}

	isRoot := isRoot()
	goCmd := "go"
	if !isRoot {
		warnOnMissingPath(config)
	} else {
		goCmd = path.Join(config.GoenvRootDirectory, "bin", goCmd)
	}

	output, err := exec.Command(goCmd, "version").Output()
	if err != nil {
		pkg.Debug(err.Error())
		return err
	}

	pkg.Success(fmt.Sprintf("Using %s", strings.TrimSuffix(string(output), "\n")))
	return nil
}

func link(config *pkg.Config, goVersion *semver.Version) error {
	// Remove the old symlink
	if _, err := os.Stat(config.GoenvRootDirectory); err == nil {
		if err = os.Remove(config.GoenvRootDirectory); err != nil {
			return fmt.Errorf("could not remove %s: %w", config.GoenvRootDirectory, err)
		}
	}

	goInstallation := path.Join(config.GoenvInstallDirectory, goVersion.Original())
	if _, err := os.Stat(goInstallation); err != nil {
		pkg.Debug(err.Error())
		return fmt.Errorf("could not find go version %s. goenv install %s", goVersion.Original(), goVersion.Original())
	}

	if err := os.Symlink(goInstallation, config.GoenvRootDirectory); err != nil {
		return fmt.Errorf("could not link: %w", err)
	}

	return nil
}

func isRoot() bool {
	return os.Geteuid() == 0
}
