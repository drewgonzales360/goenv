package pkg

import (
	"fmt"
	"os"
)

const (
	colorRed   string = "\033[31m"
	colorGreen string = "\033[32m"
	colorReset string = "\033[0m"
)

func Success(mesg string) {
	fmt.Printf("😎 %s%s%s\n", colorGreen, mesg, colorReset)
}

func Fail(mesg string) {
	fmt.Printf("😭 %s%s%s\n", colorRed, mesg, colorReset)
}

func Debug(mesg string) {
	logLevel := os.Getenv("GOENV_LOG")
	if logLevel == "DEBUG" {
		fmt.Printf("🤔 %s%s%s\n", colorRed, mesg, colorReset)
	}
}
