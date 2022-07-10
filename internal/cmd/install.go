package cmd

import (
	"fmt"
	"path"

	"github.com/Masterminds/semver"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/drewgonzales360/goenv/internal/pkg"
)

func InstallCommand(c *cli.Context) error {
	version, err := parseVersionArg(c)
	if err != nil {
		return err
	}

	config, err := parseConfig(c)
	if err != nil {
		return err
	}

	if err := install(config, version); err != nil {
		return err
	}
	return nil
}

func install(config *pkg.Config, version string) error {
	if inaccessible := pkg.CheckRW(config); len(inaccessible) > 0 {
		return fmt.Errorf(PermError, inaccessible)
	}

	goVersion, err := semver.NewVersion(version)
	if err != nil {
		return fmt.Errorf("could not parse version as a semver")
	}

	filePath, err := pkg.DownloadFile(goVersion)
	if err != nil {
		return errors.Wrap(err, "could not download go")
	}

	err = pkg.ExtractTarGz(filePath, path.Join(config.GoenvInstallDirectory, goVersion.Original()))
	if err != nil {
		return errors.Wrap(err, "could not extract go")
	}

	return use(config, version)
}
