// https://www.codewars.com/kata/is-my-friend-cheating
package kata

import "math"

func RemovNb(m uint64) [][2]uint64 {
	var numbers [][2]uint64
	s := (m * (m + 1)) / 2
	for a := uint64(1); a <= m; a++ {
		b := float64(s-a) / float64(a+1.0)
		if b < float64(m) && math.Mod(b, 1) == 0 {
			numbers = append(numbers, [2]uint64{a, uint64(b)})
		}
	}
	return numbers
}
