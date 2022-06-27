package cmd

import (
	"context"
	"fmt"

	"github.com/drewgonzales360/goenv/internal/pkg"
	"github.com/urfave/cli/v2"
)

const (
	PermError string = "you do not have access to %v"
)

func parseVersionArg(c *cli.Context) (string, error) {
	if c.NArg() != 1 {
		return "", fmt.Errorf("this command only accepts one parameter")
	}
	return c.Args().First(), nil
}

func parseConfig(c *cli.Context) (*pkg.Config, error) {
	config, ok := c.Context.Value("config").(*pkg.Config)
	if !ok {
		return nil, fmt.Errorf("could not create config")
	}
	return config, nil
}

func BeforeActionParseConfig(c *cli.Context) error {
	c.Context = context.WithValue(c.Context, "config", pkg.ReadConfig())
	pkg.Debug(fmt.Sprintf("%+v", pkg.ReadConfig()))
	return nil
}
