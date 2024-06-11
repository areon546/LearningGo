package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord     = "Go!"
	numberOfLoops = 3
	write         = "write"
	sleep         = "sleep"
)

func Countdown(out io.Writer, spy Sleeper) {
	for i := numberOfLoops; i > 0; i-- {
		output(out, i)
		spy.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func output(out io.Writer, a any) {
	fmt.Fprintln(out, a)
}

// Sleeper interface
type Sleeper interface {
	Sleep()
}

// CountDownOps
type CountDownOps struct {
	Calls      []string
	sleepCalls int
}

func (c *CountDownOps) Write([]byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

func (c *CountDownOps) Sleep() {
	c.Calls = append(c.Calls, sleep)
	c.sleepCalls++
}

// ConfigSleeper
type ConfigSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigSleeper) Sleep() {
	c.sleep(c.duration)
}

// SpyTime
type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
