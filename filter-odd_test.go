package filterOdd

import (
	"reflect"
	"testing"
)

func TestFilterOdd(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Test 1",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			want: []int{1, 3, 5, 7},
		},
		{
			name: "Test 2",
			nums: []int{2, 4, 6, 8, 10},
			want: []int{},
		},
		{
			name: "Test 3",
			nums: []int{0, 1, 2, 3, 4, 5},
			want: []int{1, 3, 5},
		},
		{
			name: "Test 4",
			nums: []int{-1, -2, -3, -4, -5},
			want: []int{-1, -3, -5},
		},
		{
			name: "Test 5",
			nums: []int{},
			want: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name, func(t *testing.T) {
				actual := filterOdd(tc.nums)
				if !reflect.DeepEqual(actual, tc.want) {
					t.Errorf("Expected %v, got %v", tc.want, actual)
				}
			},
		)
	}
}
