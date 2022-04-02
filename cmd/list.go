package cmd

import (
	"fmt"
	"os"

	"github.com/drewgonzales360/goenv/pkg"
	"github.com/urfave/cli/v2"
)

func ListCommand(c *cli.Context) error {
	versions, err := os.ReadDir(InstallDirectory)
	if err != nil {
		pkg.Fail(err.Error())
	}

	for _, version := range versions {
		fmt.Println(version.Name())
	}

	return nil
}
