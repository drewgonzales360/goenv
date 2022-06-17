package pkg

import (
	"os"

	"github.com/fatih/color"
)

func Success(mesg string) {
	color.Green("😎 %s", mesg)
}

func Fail(mesg string) {
	color.Red("😭 %s", mesg)
}

func Debug(mesg string) {
	logLevel := os.Getenv("GOENV_LOG")
	if logLevel == "DEBUG" {
		color.Blue("🤔 %s", mesg)
	}
}
