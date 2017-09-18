//https://www.codewars.com/kata/going-to-zero-or-to-infinity/train/go

package kata

import "math"

func Going(n int) float64 {
	a := 1.0
	b := 1.0

	for c := 2; c <= n; c++ {
		b *= float64(c)
		a += b

		for a > 1e6 || b > 1e6 {
			b /= 10.0
			a /= 10.0
		}
	}

	result := math.Trunc(float64(a)/float64(b)*1e6) / 1e6
	return result
}
