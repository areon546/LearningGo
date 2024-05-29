package main

import "fmt"

func main() {
	fmt.Println(Hello("world"))
}

const engHelloPref = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return engHelloPref + name
}
