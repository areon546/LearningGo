package main

import (
	"fmt"

	"github.com/areon546/LearningGo/goFiles/concurrency"
)

func main() {
	// concurrencyF()
}

func concurrencyF() {
	urls := []string{"1", "2"}
	concurrency.LoopWebsites(urls)
}

func mockingF() {
	// durationA := 1 * time.Second
	// sleeper := &mocking.ConfigSleeper{durationA, time.Sleep}
	// mocking.Countdown(os.Stdout, sleeper)
	return
}

func LoopWebsites(urls []string) {
	for index, value := range urls {
		fmt.Println(index, value)
	}
}
