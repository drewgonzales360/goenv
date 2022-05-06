package pkg

import (
	"os"
	"time"
)

func CheckRW() []string {
	accessDenied := []string{}
	currentTime := time.Now().Local()

	dirs := []string{"/usr/local/goenv", "/usr/local/go"}
	for _, dir := range dirs {
		if _, err := os.Stat(dir); err == nil {
			if err := os.Chtimes(dir, currentTime, currentTime); err != nil {
				Debug(err.Error())
				accessDenied = append(accessDenied, dir)
			}
		} else {
			if err := os.MkdirAll(dir, 0755); err != nil {
				Debug(err.Error())
				accessDenied = append(accessDenied, dir)
			}
		}
	}

	return accessDenied
}
