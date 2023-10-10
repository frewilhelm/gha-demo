package lib

import "testing"

func TestPower(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{a: 2, b: 1, expected: 2},
		{a: 10, b: 2, expected: 100},
		{a: 12, b: 6, expected: 2985984},
		{a: 15, b: 3, expected: 3375},
		{a: 36, b: 12, expected: 1679616},
	}

	for _, tc := range testCases {
		result := Power(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Got %d - %d, = %d, want %d", tc.a, tc.b, tc.expected, result)
		}
	}
}
