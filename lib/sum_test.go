package lib

import "testing"

func TestSum(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{a: 1, b: 2, expected: 4},
		{a: 11, b: 2, expected: 13},
		{a: 1, b: 21, expected: 26},
		{a: 15, b: 15, expected: 30},
		{a: 31, b: 12, expected: 43},
	}

	for _, tc := range testCases {
		result := Sum(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Sum %d, %d, = %d, want %d", tc.a, tc.b, tc.expected, result)
		}
	}
}
