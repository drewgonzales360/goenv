package cmd

import (
	"fmt"
	"os"

	"github.com/Masterminds/semver"
	"github.com/urfave/cli/v2"

	"github.com/drewgonzales360/goenv/pkg"
)

const (
	InstallDirectory string = "/usr/local/go/"
)

func Install(c *cli.Context) error {
	version := ""
	if c.NArg() > 0 {
		version = c.Args().Get(0)
	}

	goVersion, err := semver.NewVersion(version)
	if err != nil {
		return fmt.Errorf("could not parse version as a semver")
	}

	downloadURL := pkg.FormatDownloadURL(*goVersion)
	filePath, err := pkg.DownloadFile(downloadURL)
	if err != nil {
		return fmt.Errorf("could not download go")
	}

	err = pkg.ExtractTarGz(filePath, InstallDirectory+goVersion.Original())
	if err != nil {
		return fmt.Errorf("could not extract go")
	}

	if err := os.Symlink("/usr/local/bin/go", InstallDirectory+goVersion.Original()+"/bin/go"); err != nil {
		return fmt.Errorf("could not link go")
	}

	return nil
}
