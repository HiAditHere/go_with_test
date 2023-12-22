package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("Countdown test 3,2,1 Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()

		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("Got %q Wanted %q", got, want)
		}

	})

	t.Run("Order of sleep and write", func(t *testing.T) {
		spyOrder := &SpyOrder{}

		Countdown(spyOrder, spyOrder)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spyOrder.calls) {
			t.Errorf("got calls %v wanted calls %v", spyOrder.calls, want)
		}
	})
}

func TestConfigureSleeper(t *testing.T) {

	sleepTime := 5 * time.Second

	spyTime := &SpyTimer{}

	sleeper := &ConfigurableSleeper{sleepTime, spyTime.Sleep}

	sleeper.Sleep()

	if spyTime.duration != sleepTime {
		t.Errorf("Wanted %v got %v", sleepTime, spyTime.duration)
	}

}
