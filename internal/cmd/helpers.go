package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/drewgonzales360/goenv/internal/pkg"
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

func warnOnMissingPath(config *pkg.Config) {
	bin := config.GoenvRootDirectory + "/bin"
	// TODO: if root installs, and root doesn't have this in the path, it'll warn
	// unnessecarily
	if path := os.Getenv("PATH"); !strings.Contains(path, bin) {
		pkg.Info(fmt.Sprintf("%s is not in your PATH", bin))
		pkg.Info(fmt.Sprintf("export PATH=%s:$PATH # to include it", bin))
	}
}
