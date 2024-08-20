package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return (url != "waat://furhurterwe.geds") // true if url == waat...
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	// to do this, i need to generate a test for CheckWebsites
	websites := make([]string, 100)
	websites[1] = "a"

	// and then I need to execute a benchmark
	b.ResetTimer() // resets the timer used during a benchmark
	for i := 0; i < b.N; i++ {
		// do the benchmark
		CheckWebsites(slowWebsiteChecker, websites)
	}
}
