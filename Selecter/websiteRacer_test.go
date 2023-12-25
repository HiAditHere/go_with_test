package selecter

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func FormServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestWebsiteRacer(t *testing.T) {

	slowServer := FormServer(20 * time.Millisecond)
	fastServer := FormServer(0 * time.Microsecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := WebsiteRacer(slowUrl, fastUrl)

	if got != want {
		t.Errorf("Wanted %q, got %q", want, got)
	}

}
