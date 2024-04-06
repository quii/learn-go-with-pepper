package arrays_and_slices

import (
	. "github.com/quii/pepper"
	. "github.com/quii/pepper/matchers/comparable"
	"github.com/quii/pepper/matchers/slice"
	"testing"
)

func TestSum(t *testing.T) {
	Expect(t, Sum([]int{1, 2, 3, 4, 5})).To(Equal(15))
	Expect(t, Sum([]int{1, 2, 3})).To(Equal(6))
}

func TestSumAll(t *testing.T) {
	Expect(
		t,
		SumAll([]int{1, 2}, []int{0, 9}),
	).To(slice.ShallowEquals([]int{3, 9}))
}
