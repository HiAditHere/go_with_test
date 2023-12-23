package concurrency

import (
	"reflect"
	"testing"
)

func mockChecker(url string) bool {
	if url == "www.something.notapplicable" {
		return false
	}

	return true
}

func TestWebsiteChecker(t *testing.T) {

	urls := []string{
		"www.abc.com",
		"www.xyz.com",
		"www.something.notapplicable",
	}

	want := map[string]bool{
		"www.abc.com":                 true,
		"www.xyz.com":                 true,
		"www.something.notapplicable": false,
	}

	got := CheckWebsites(mockChecker, urls)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
