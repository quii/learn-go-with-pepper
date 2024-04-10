package concurrency

import (
	"github.com/quii/pepper"
	. "github.com/quii/pepper/matchers/comparable"
	"github.com/quii/pepper/matchers/maps"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	pepper.Expect(t, CheckWebsites(mockWebsiteChecker, websites)).To(
		maps.HaveKey("http://google.com", Equal(true)),
		maps.HaveKey("http://blog.gypsydave5.com", Equal(true)),
		maps.HaveKey("waat://furhurterwe.geds", Equal(false)),
	)
}
