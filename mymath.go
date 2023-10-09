package mymath

import (
	"math"
)

// Sqrt - функция для вычисления квадратного корня из числа
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}
func Abs(x float64) float64 {
	return math.Abs(x)
}
func Max(x, y float64) float64 {
	return math.Max(x, y)
}
func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}
