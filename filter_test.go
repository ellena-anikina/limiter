package filter

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	testCases := []struct {
		name      string
		predicate func(int) bool
		iterable  []int
		want      []int
	}{
		{
			name:      "even numbers",
			predicate: func(x int) bool { return x%2 == 0 },
			iterable:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:      []int{2, 4, 6, 8, 10},
		},
		{
			name:      "odd numbers",
			predicate: func(x int) bool { return x%2 != 0 },
			iterable:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:      []int{1, 3, 5, 7, 9},
		},
		{
			name:      "numbers greater than 5",
			predicate: func(x int) bool { return x > 5 },
			iterable:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:      []int{6, 7, 8, 9, 10},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := filter(tc.predicate, tc.iterable)
			if fmt.Sprintf("%v", actual) != fmt.Sprintf("%v", tc.want) {
				t.Errorf("Expected %v, got %v", tc.want, actual)
			}
		})
	}
}
