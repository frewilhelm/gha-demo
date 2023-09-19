package lib

import "testing"

func TestSub(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{a: 1, b: 2, expected: -1},
		{a: 11, b: 2, expected: 9},
		{a: 1, b: 21, expected: -20},
		{a: 15, b: 15, expected: 0},
		{a: 31, b: 12, expected: 19},
	}

	for _, tc := range testCases {
		result := Sub(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Got %d - %d, = %d, want %d", tc.a, tc.b, tc.expected, result)
		}
	}
}
