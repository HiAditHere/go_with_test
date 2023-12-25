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

	t.Run("Test for one website being faster than the other", func(t *testing.T) {
		slowServer := FormServer(20 * time.Millisecond)
		fastServer := FormServer(0 * time.Microsecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := EnhancedWebsiteRacer(slowUrl, fastUrl, 10*time.Second)

		if err != nil {
			t.Fatalf("did not expect any error but got one %v", err)
		}

		if got != want {
			t.Errorf("Wanted %q, got %q", want, got)
		}
	})

	t.Run("Timeout incase of more than certain time", func(t *testing.T) {
		server := FormServer(25 * time.Millisecond)

		defer server.Close()

		_, err := EnhancedWebsiteRacer(server.URL, server.URL, time.Duration(20*time.Millisecond))

		if err == nil {
			t.Errorf("wanted an error but got none")
		}
	})

}
