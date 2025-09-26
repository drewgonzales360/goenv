package cmd

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/drewgonzales360/goenv/internal/pkg"
	"github.com/spf13/cobra"
)

func ScriptCommand(cmd *cobra.Command, args []string) error {
	version := args[0]
	config := pkg.ReadConfig()

	global, err := cmd.Flags().GetBool("global")
	if err != nil {
		return err
	}

	goVersion, err := semver.NewVersion(version)
	if err != nil {
		return fmt.Errorf("could not parse %s: %w", version, err)
	}

	rootDir := config.GoenvRootDirectory
	installDir := config.VersionInstallPath(goVersion)
	if global {
		rootDir = "/usr/local/go"
		installDir = "/usr/local/go"
	}

	return printInstallScript(rootDir, installDir, version)
}

func printInstallScript(rootDir , installDir, version string) error {
	goVersion, err := semver.NewVersion(version)
	if err != nil {
		return fmt.Errorf("could not parse %s: %w", version, err)
	}

	fmt.Println(pkg.InstallScript(rootDir, installDir, goVersion))

	return nil
}
