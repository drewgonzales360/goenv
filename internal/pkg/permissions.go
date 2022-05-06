package pkg

import (
	"fmt"
	"os"
	"time"
)

func CheckRW() []string {
	accessDenied := []string{}
	currentTime := time.Now().Local()

	installDirectory := "/usr/local/goenv"
	if err := os.MkdirAll(installDirectory, 0755); err != nil {
		Debug(err.Error())
		accessDenied = append(accessDenied, installDirectory)
	}
	if err := os.Chtimes(installDirectory, currentTime, currentTime); err != nil {
		Debug(err.Error())
		accessDenied = append(accessDenied, installDirectory)
	}
	Debug(fmt.Sprintf("Created %s", installDirectory))

	usrLocalGo := "/usr/local/go"
	if _, err := os.Create(usrLocalGo); err != nil {
		Debug(err.Error())
		accessDenied = append(accessDenied, usrLocalGo)
	}
	if err := os.Chtimes(usrLocalGo, currentTime, currentTime); err != nil {
		Debug(err.Error())
		accessDenied = append(accessDenied, usrLocalGo)
	}
	Debug(fmt.Sprintf("Created %s", usrLocalGo))

	return accessDenied
}
