package lib

import "testing"

func TestMult(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{a: 1, b: 2, expected: 2},
		{a: 11, b: 2, expected: 22},
		{a: 1, b: 21, expected: 21},
		{a: 15, b: 15, expected: 225},
		{a: 31, b: 12, expected: 372},
	}

	for _, tc := range testCases {
		result := Mult(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Got %d - %d, = %d, want %d", tc.a, tc.b, tc.expected, result)
		}
	}
}
