package concurrency

import "fmt"

type WebsiteChecker func(string) bool

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	result := make(map[string]bool)

	// iterates through urls and stores the result of wc[url] in result
	for a, b := range urls {
		fmt.Println(a, b)
	}

	return result
}

func LoopWebsites(urls []string) {
	for a, b := range urls {
		fmt.Println(a, b)
	}
}
