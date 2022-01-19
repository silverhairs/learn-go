package selection

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(delay)
		res.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	fast := "https://fast.com"
	slow := "https://slow.com"

	t.Run("returns the url with the fastest request", func(t *testing.T) {
		fastServer := makeDelayedServer(0 * time.Microsecond)
		slowServer := makeDelayedServer(0 * time.Millisecond)

		defer fastServer.Close()
		defer slowServer.Close()

		expect := fast
		found := Racer(fast, slow)

		if expect != found {
			t.Errorf("expected '%s' but found '%s'", expect, found)
		}
	})

}
