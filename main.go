package main

import (
	"os"
	"time"

	"github.com/areon546/LearningGo/goFiles/mocking"
)

func main() {
	durationA := 1 * time.Second
	sleeper := &mocking.ConfigSleeper{durationA, time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)

	// var err error // initial value is nil
	// fmt.Println(err.Error()) //

	//
	// input := "Test sentence"
	// fmt.Println(strings.Fields(input)[0])

	// fmt.Println(Hello("world", ""))
}
