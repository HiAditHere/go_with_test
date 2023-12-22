package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const startCountDown = 3
const finalWord = "Go!"
const sleep = "sleep"
const write = "write"

type SpySleeper struct {
	calls int
}

type Sleeper interface {
	Sleep()
}

func (s *SpySleeper) Sleep() {
	s.calls++
}

type SpyOrder struct {
	calls []string
}

func (s *SpyOrder) Sleep() {
	s.calls = append(s.calls, sleep)
}

func (s *SpyOrder) Write(p []byte) (n int, err error) {
	s.calls = append(s.calls, write)
	return
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

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTimer struct {
	duration time.Duration
}

func (s *SpyTimer) Sleep(time time.Duration) {
	s.duration = time
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	//sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
