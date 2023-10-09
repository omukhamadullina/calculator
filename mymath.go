package mymath

import (
	"math"
)

func Sqrt(x float64) int {
	return int(math.Sqrt(x))
}
func Abs(x float64) int {
	return int(math.Abs(x))
}
func Max(x, y float64) int {
	return int(math.Max(x, y))
}
func Yn(n int, x float64) int{
	return int(math.Yn(n, x))
}
