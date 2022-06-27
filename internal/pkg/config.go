package pkg

import "os"

type Config struct {
	// This is the directory holding multiple Go installations is.
	// GOENV_ROOT_DIR
	GoenvRootDirectory string

	// Install directory defaults to /usr/local/go and can be configured with
	// GOENV_INSTALL_DIR.
	GoenvInstallDirectory string
}

const (
	DefaultGoenvRootDirectory = "/usr/local/goenv"
	DefaultGoInstallDirectory = "/usr/local/go"
)

func ReadConfig() *Config {
	rootDir := os.Getenv("GOENV_ROOT_DIR")
	if rootDir == "" {
		rootDir = DefaultGoenvRootDirectory
	}

	installDir := os.Getenv("GOENV_INSTALL_DIR")
	if installDir == "" {
		installDir = DefaultGoInstallDirectory
	}

	return &Config{
		rootDir,
		installDir,
	}
}
