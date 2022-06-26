package pkg_test

import (
	"testing"

	"github.com/Masterminds/semver"
	"github.com/drewgonzales360/goenv/internal/pkg"
)

// TestShasumsAreAccurate will test that the sums that we put in
// are correct. If we haven't added the shasums for new versions
// we won't fail on those so as not to block new versions from being
// downloaded.
func TestShasumsAreAccurate(t *testing.T) {
	for _, v := range pkg.GoVersions {
		for version := range v {
			_, err := pkg.DownloadFile(*semver.MustParse(version))
			if err != nil {
				t.Logf("couldn't download %s", version)
				t.Fail()
			}
		}
	}
}
