package pkg

import (
	"bytes"
	"fmt"
	"os"
	"text/tabwriter"
)

type Config struct {
	// The file placed at GoenvRootDirectory is a symlink to the GoenvInstallDirectory
	GoenvRootDirectory string

	// Install directory defaults to /usr/local/goenv and can be configured with
	GoenvInstallDirectory string
}

const (
	DefaultGoenvRootDirectory = "/usr/local/go"
	DefaultGoInstallDirectory = "/usr/local/goenv"

	// This should also be the users GOROOT.
	GoEnvRootDirEnvVar    = "GOENV_ROOT_DIR"
	GoEnvInstallDirEnvVar = "GOENV_INSTALL_DIR"
)

// ReadConfig reads the environment variables for a user and creates
// a config. If we need any additional config, it'll be parsed in here.
func ReadConfig() *Config {
	rootDir := os.Getenv(GoEnvRootDirEnvVar)
	if rootDir == "" {
		rootDir = DefaultGoenvRootDirectory
	}

	installDir := os.Getenv(GoEnvInstallDirEnvVar)
	if installDir == "" {
		installDir = DefaultGoInstallDirectory
	}

	return &Config{
		rootDir,
		installDir,
	}
}

func (c *Config) String() string {
	buf := bytes.Buffer{}
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', 0)

	rootDir := fmt.Sprintf("%s:\t%s", GoEnvRootDirEnvVar, c.GoenvRootDirectory)
	if c.GoenvRootDirectory == DefaultGoenvRootDirectory {
		rootDir += "\t(default)\n"
	} else {
		rootDir += "\t(set by environment variable)\n"
	}
	w.Write([]byte(rootDir))

	installDir := fmt.Sprintf("%s:\t%s", GoEnvInstallDirEnvVar, c.GoenvInstallDirectory)
	if c.GoenvInstallDirectory == DefaultGoInstallDirectory {
		installDir += "\t(default)\n"
	} else {
		installDir += "\t(set by environment variable)\n"
	}
	w.Write([]byte(installDir))
	w.Flush()
	return buf.String()
}
