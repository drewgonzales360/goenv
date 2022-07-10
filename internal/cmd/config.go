package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func ConfigCommand(c *cli.Context) error {
	config, err := parseConfig(c)
	if err != nil {
		return err
	}

	fmt.Println(config)
	warnOnMissingPath(config)
	return nil
}
