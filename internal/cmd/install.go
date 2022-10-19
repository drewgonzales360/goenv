package cmd

import (
	"fmt"
	"path"

	"github.com/Masterminds/semver"
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
		return fmt.Errorf("could not parse %s: %w", version, err)
	}

	filePath, err := pkg.DownloadFile(goVersion)
	if err != nil {
		return fmt.Errorf("could not download go: %w", err)
	}

	err = pkg.ExtractTarGz(filePath, path.Join(config.GoenvInstallDirectory, goVersion.Original()))
	if err != nil {
		return fmt.Errorf("could not extract go: %w", err)
	}

	err = use(config, version)
	if err != nil {
		return err
	}

	return nil
}
