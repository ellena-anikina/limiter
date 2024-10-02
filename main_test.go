package merge_channels

import (
	"reflect"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 2
	ch3 <- 3

	close(ch1)
	close(ch2)
	close(ch3)

	merged := merge(ch1, ch2, ch3)

	results := make([]int, 0, 3)
	for i := range merged {
		results = append(results, i)
	}

	sort.Ints(results)

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected %v, got %v", expected, results)
	}
}
