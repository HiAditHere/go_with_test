package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const startCountDown = 3
const finalWord = "Go!"

type SpySleeper struct {
	calls int
}

type Sleeper interface {
	Sleep()
}

func (s *SpySleeper) Sleep() {
	s.calls++
}

func Countdown(out io.Writer, sleeper Sleeper) {

	for i := startCountDown; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)

}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
