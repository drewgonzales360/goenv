package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
)

func ConfigCommand(c *cli.Context) error {
	config, err := parseConfig(c)
	if err != nil {
		return err
	}

	stringConfig := fmt.Sprintf("%+v\n", *config)
	re, err := regexp.Compile(`[&\{\}]`)
	if err != nil {
		return err
	}

	sanitizedConfig := re.ReplaceAllString(stringConfig, "")
	listConfigs := strings.Fields(sanitizedConfig)
	for _, l := range listConfigs {
		fmt.Println(l)
	}

	return nil
}
