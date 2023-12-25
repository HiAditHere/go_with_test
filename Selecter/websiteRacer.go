package selecter

import (
	"fmt"
	"net/http"
	"time"
)

func MeasureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	Duration := time.Since(start)

	return Duration

}

func WebsiteRacer(a, b string) string {

	aDuration := MeasureResponseTime(a)
	bDuration := MeasureResponseTime(b)

	if aDuration > bDuration {
		return b
	}

	return a
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

func EnhancedWebsiteRacer(a, b string, duration time.Duration) (url string, err error) {

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(duration):
		return "", fmt.Errorf("website %v and %v timedout", a, b)
	}
}
