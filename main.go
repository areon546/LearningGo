package main

import (
	"os"

	"github.com/areon546/LearningGo/goFiles/injection"
)

func main() {
	injection.Greet(os.Stdout, "Chris")

	// var err error // initial value is nil
	// fmt.Println(err.Error()) //

	//
	// input := "Test sentence"
	// fmt.Println(strings.Fields(input)[0])

	// fmt.Println(Hello("world", ""))
}
