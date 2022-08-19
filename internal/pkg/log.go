package pkg

import (
	"os"

	"github.com/fatih/color"
)

func Success(mesg string) {
	color.Green("😎 %s", mesg)
}

func Error(mesg string) {
	color.Red("😭 %s", mesg)
}

func Info(mesg string) {
	color.White("😃 %s", mesg)
}

func Warn(mesg string) {
	color.Yellow("😥 %s", mesg)
}

func Debug(mesg string) {
	logLevel := os.Getenv("GOENV_LOG")
	if logLevel == "DEBUG" {
		color.Blue("🤔 %s", mesg)
	}
}
