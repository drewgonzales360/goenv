package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"

	"github.com/drewgonzales360/goenv/internal/pkg"
)

func ListCommand(c *cli.Context) error {
	config, err := parseConfig(c)
	if err != nil {
		return err
	}

	versions, err := os.ReadDir(config.GoenvRootDirectory)
	if err != nil {
		return err
	}

	names := []string{}
	for _, version := range versions {
		names = append(names, version.Name())
	}
	installed := pkg.CreateGoVersionList(names)
	color.New(color.FgCyan, color.Bold).Println("Installed Versions:")
	pkg.Print(&installed)

	if c.Bool("available") {
		color.New(color.FgCyan, color.Bold).Println("Available Versions:")
		pkg.Print(&pkg.GoVersions)
	}

	return nil
}
