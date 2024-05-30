package main

import "fmt"

func main() {
	fmt.Println(Hello("world", ""))
}

const (
	engHelloPref = "Hello, "

	spanish      = "Spanish"
	spnHelloPref = "Hola, "

	french       = "French"
	freHelloPref = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return checkGreetingPreffix(language) + name
}

func checkGreetingPreffix(language string) (preffix string) {
	switch language {
	default:
		preffix = engHelloPref
	case spanish:
		preffix = spnHelloPref
	case french:
		preffix = freHelloPref
	}

	return
}
