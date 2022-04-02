package pkg

import (
	"os"
	"time"
)

func CheckRW() []string {
	accessDenied := []string{}
	currentTime := time.Now().Local()

	installDirectory := "/usr/local/"
	if err := os.Chtimes(installDirectory, currentTime, currentTime); err != nil {
		Debug(err.Error())
		accessDenied = append(accessDenied, installDirectory)
	}

	binDirectory := "/usr/local/bin/"
	if err := os.Chtimes(binDirectory, currentTime, currentTime); err != nil {
		Debug(err.Error())
		accessDenied = append(accessDenied, binDirectory)
	}

	return accessDenied
}
