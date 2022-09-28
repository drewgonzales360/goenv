package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/drewgonzales360/goenv/internal/pkg"
)

type GoEnvContextKey string

const (
	PermError            string          = "you do not have access to %v"
	configContextKey     GoEnvContextKey = "config"
	RootDirContextKey    string          = "root-dir"
	InstallDirContextKey string          = "install-dir"
)

func parseVersionArg(c *cli.Context) (string, error) {
	if c.NArg() != 1 {
		return "", fmt.Errorf("this command only accepts one parameter")
	}
	return c.Args().First(), nil
}

func parseConfig(c *cli.Context) (*pkg.Config, error) {
	config, ok := c.Context.Value(configContextKey).(*pkg.Config)
	if !ok {
		return nil, fmt.Errorf("could not create config")
	}
	return config, nil
}

func BeforeActionParseConfig(ctx *cli.Context) error {
	config := &pkg.Config{
		GoenvInstallDirectory: ctx.Path(InstallDirContextKey),
		GoenvRootDirectory:    ctx.Path(RootDirContextKey),
	}
	pkg.Debug(fmt.Sprintf("Config:\n%+v", config))
	ctx.Context = context.WithValue(ctx.Context, configContextKey, config)
	return nil
}

func warnOnMissingPath(config *pkg.Config) {
	bin := config.GoenvRootDirectory + "/bin"
	if path := os.Getenv("PATH"); !strings.Contains(path, bin) {
		pkg.Info(fmt.Sprintf("%s is not in your PATH", bin))
		pkg.Info(fmt.Sprintf("export PATH=%s:$PATH", bin))
	}
}
