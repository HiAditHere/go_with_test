package concurrency

import (
	"reflect"
	"testing"
	"time"
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

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Microsecond)
	return true
}

func BenchmarkWebsiteChecker(b *testing.B) {
	urls := make([]string, 100)

	for i := 1; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()

	for i := 1; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
