// https://www.codewars.com/kata/check-if-a-triangle-is-an-equable-triangle/train/go

package kata

import "math"

func EquableTriangle(a, b, c int) bool {
	perimeter := a + b + c
	area := CalculateArea(float64(a), float64(b), float64(c))
	return perimeter == area
}

func CalculateArea(a, b, c float64) int {
	s := (a + b + c) / 2
	area := math.Sqrt(s * (s - a) * (s - b) * (s - c))
	return int(area)
}
