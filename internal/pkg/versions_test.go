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
	t.Skip("skipping download test because it's slow")

	for _, v := range pkg.GoVersions {
		for version := range v {
			v := semver.MustParse(version)
			// If we don't know the hash, don't test the download.
			if pkg.GetHash(v) == "" {
				continue
			}

			_, err := pkg.DownloadFile(v)
			if err != nil {
				t.Logf("couldn't download %s: %s", version, err.Error())
				t.Fail()
			}
		}
	}
}
