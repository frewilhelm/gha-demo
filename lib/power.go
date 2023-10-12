package lib

import "math"

func Power(a int, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
