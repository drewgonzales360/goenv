package pkg_test

import (
	"testing"

	"github.com/drewgonzales360/goenv/internal/pkg"
)

func TestDiscover(t *testing.T) {
	releases, err := pkg.GetGoVersions(false)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if len(releases) == 0 {
		t.Fail()
	}
	t.Logf("%+v", releases)
}
