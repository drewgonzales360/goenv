// ///////////////////////////////////////////////////////////////////////
// Copyright 2022 Drew Gonzales
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ///////////////////////////////////////////////////////////////////////
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
	PermError string          = "you do not have access to %v"
	config    GoEnvContextKey = "config"
)

// parseVersionArg ensures that the subcommands that accept a parameter
// only spcify one parameter.
func parseVersionArg(c *cli.Context) (string, error) {
	if c.NArg() != 1 {
		return "", fmt.Errorf("this command only accepts one parameter")
	}
	return c.Args().First(), nil
}

// parseConfig reads the config from the context, assuming that BeforeActionParseConfig
// puts the config *in* the context.
func parseConfig(c *cli.Context) (*pkg.Config, error) {
	config, ok := c.Context.Value(config).(*pkg.Config)
	if !ok {
		return nil, fmt.Errorf("could not create config")
	}
	return config, nil
}

// BeforeActionParseConfig adds the config to the context so it can be read
// by parseConfig.
func BeforeActionParseConfig(c *cli.Context) error {
	c.Context = context.WithValue(c.Context, config, pkg.ReadConfig())
	return nil
}

// AfterAction checks for new versions of Goenv and Go.
func AfterAction(c *cli.Context) error {
	pkg.CheckLatestGoenv(c.App.Version)
	pkg.CheckLatestGo()
	return nil
}

// warnOnMissingPath does a best effort to let you know that Go can't be called.
// This will sometimes warn unnecessarily if the user runs as root.
func warnOnMissingPath(config *pkg.Config) {
	bin := config.GoenvRootDirectory + "/bin"
	if path := os.Getenv("PATH"); !strings.Contains(path, bin) {
		pkg.Info(fmt.Sprintf("%s is not in your PATH", bin))
		pkg.Info(fmt.Sprintf("export PATH=%s:$PATH", bin))
	}
}
