package pkg

import (
	"testing"
)

func TestDiscover(t *testing.T) {
	releases, err := getGoVersions(false)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if len(releases) == 0 {
		t.Fail()
	}
	t.Logf("%+v", releases)
}
