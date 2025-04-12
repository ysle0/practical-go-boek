package pkgquery

import (
	"testing"

	"github.com/ysle0/shared"
)

func TestFetchPkgData(t *testing.T) {
	json := `[{"name": "pkg1", "version": "1.0.0"}, {"name": "pkg2", "version": "2.0.0"}]`
	ts := shared.StartTestHttpServerWithJson(json)
	defer ts.Close()

	actual, err := fetchPkgData(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	if len(actual) != 2 {
		t.Fatalf("expect: 2 packages, got %d", len(actual))
	}
}
