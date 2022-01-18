package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "test://testing.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"http://example.com",
		"test://testing.com",
	}

	got := CheckWebsites(mockWebsiteChecker, websites)
	expect := map[string]bool{
		"https://google.com": true,
		"http://example.com": true,
		"test://testing.com": false,
	}

	if !reflect.DeepEqual(expect, got) {
		t.Fatalf("\nexpected: %v, \nbut got: %v", expect, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, []string{
			"https://google.com",
			"http://example.com",
			"test://testing.com",
		})
	}
}
