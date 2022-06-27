package pkg

import (
	"os"
	"time"
)

func CheckRW(config *Config) []string {
	accessDenied := []string{}
	currentTime := time.Now().Local()

	dirs := []string{config.GoenvRootDirectory, config.GoenvInstallDirectory}
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
