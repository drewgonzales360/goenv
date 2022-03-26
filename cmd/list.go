package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func List(c *cli.Context) error {
	fmt.Println("Hello world!")
	return nil
}
