package pkg

import (
	"os"
	"time"
)

// CheckRW checks if the user can install and use Go by making sure
// they have read and write access to the directories where go will
// be installed.
func CheckRW(config *Config) []string {
	accessDenied := []string{}
	currentTime := time.Now().Local()

	if err := removeDeadLink(config.GoenvRootDirectory); err != nil {
		Debug(err.Error())
	}

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

// removeDeadLink checks if the go root is pointed at an uninstalled
// go version. If that's the case, then we remove the link.
func removeDeadLink(path string) error {
	installDir, err := os.Readlink(path)
	if err != nil {
		return err
	}

	_, err = os.Stat(installDir)
	if err == nil {
		return nil
	}

	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}
