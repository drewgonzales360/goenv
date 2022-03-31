package version

import (
	"github.com/Masterminds/semver"
)

// AppName represents the name of the application
const AppName = "goenv"

// Version semvers the app
var Semver string = "unknown; please create an issue for the maintainers"

func Version() *semver.Version {
	return semver.MustParse(Semver)
}
