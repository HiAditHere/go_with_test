package mocking

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("Countdown test", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer)

		got := buffer.String()

		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("Got %q Wanted %q", got, want)
		}

	})
}
