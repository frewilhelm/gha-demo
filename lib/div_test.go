package lib

import "testing"

func TestDiv(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{a: 2, b: 1, expected: 2},
		{a: 10, b: 2, expected: 5},
		{a: 12, b: 6, expected: 2},
		{a: 15, b: 15, expected: 1},
		{a: 36, b: 12, expected: 3},
	}

	for _, tc := range testCases {
		result := Div(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Got %d - %d, = %d, want %d", tc.a, tc.b, tc.expected, result)
		}
	}
}
