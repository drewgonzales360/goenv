package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const (
	UsrLocalBin      string = "/usr/local/bin/"
	PermError        string = "are you sure you're root? you do not have access to %v"
	InstallDirectory string = "/usr/local/go/"
)

func parseVersionArg(c *cli.Context) (string, error) {
	if c.NArg() != 1 {
		return "", fmt.Errorf("this command only accepts one parameter")
	}
	return c.Args().First(), nil
}
