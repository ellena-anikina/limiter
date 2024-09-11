package cat

import (
	"testing"
)

func TestCat(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		n    int
		want string
	}{
		{
			name: "Test 1",
			s:    "awasdeabbd",
			n:    1,
			want: "awsdeb",
		},
		{
			name: "Test 2",
			s:    "abbbacddd",
			n:    2,
			want: "abbacdd",
		},
		{
			name: "Test 3",
			s:    "abbbacddd",
			n:    3,
			want: "abbbacddd",
		},
		{
			name: "Test 4",
			s:    "dsadsdddffasau",
			n:    2,
			want: "dsadsffau",
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name, func(t *testing.T) {
				actual := cat(tc.s, tc.n)
				if actual != tc.want {
					t.Errorf("Expected %v, got %v", tc.want, actual)
				}
			},
		)
	}
}
