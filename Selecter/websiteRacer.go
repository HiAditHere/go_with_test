package selecter

import (
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
