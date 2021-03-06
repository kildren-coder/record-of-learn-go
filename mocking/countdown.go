package mocking

import (
	"fmt"
	"io"
)

type Sleeper interface {
	Sleep()
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer,sleeper Sleeper) {
	for i := countdownStart; i > 0; i -- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}