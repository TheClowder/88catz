//https://www.codewars.com/kata/playing-on-a-chessboard

package main

import "math"

func Game(n int) []int {
	if math.Mod(float64(n), 2.0) == 0 {
		return []int{n * n / 2}
	}
	return []int{n * n, 2}
}
