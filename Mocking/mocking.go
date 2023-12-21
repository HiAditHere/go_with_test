package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const startCountDown = 3
const finalWord = "Go!"

func Countdown(out io.Writer) {

	for i := startCountDown; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, finalWord)

}

func main() {
	Countdown(os.Stdout)
}
