package sync

import (
	. "github.com/quii/pepper"
	. "github.com/quii/pepper/matchers/comparable"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		Expect(t, counter).To(HaveCount(Equal(3)))
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		Expect(t, counter).To(HaveCount(Equal(wantedCount)))
	})
}

// Note how want is a Matcher[int] rather than an int. Whilst this adds wordiness
// in the asserting, it gives flexibility. If I set it to int, I could only ever
// assert the counter is equal to a value, by giving a matcher, I can assert all
// sorts of things, like GreaterThan, or IsEven, etc

func HaveCount(want Matcher[int]) Matcher[Counter] {
	return func(actual Counter) MatchResult {
		return want(actual.Value())
	}
}
