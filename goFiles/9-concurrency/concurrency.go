package concurrency

import "fmt"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	result := make(map[string]bool)

	// iterates through urls and stores the result of wc[url] in result
	for _, url := range urls {
		go func(u string) {
			result[u] = wc(u)
		}(url)
	}

	return result
}

func LoopWebsites(urls []string) {
	for index, value := range urls {
		fmt.Println(index, value)
	}
}
