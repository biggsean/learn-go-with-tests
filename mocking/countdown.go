package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord              = "Go!"
	countdownStart         = 3
	countdownSleepDuration = 1 * time.Second
)

// Sleeper interface
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper is a configurable sleeper
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep method for ConfigurableSleeper
func (c ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Countdown counts it down!
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{countdownSleepDuration, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
