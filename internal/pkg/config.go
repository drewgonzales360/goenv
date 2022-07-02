package pkg

import "os"

type Config struct {
	// This is the directory holding multiple Go installations is.
	GoenvRootDirectory string

	// Install directory defaults to /usr/local/go and can be configured with
	GoenvInstallDirectory string
}

const (
	DefaultGoenvRootDirectory = "/usr/local/goenv"
	DefaultGoInstallDirectory = "/usr/local/go"

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
