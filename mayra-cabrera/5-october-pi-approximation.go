//https://www.codewars.com/kata/pi-approximation

package kata

import (
	"fmt"
	"math"
)

func IterPi(epsilon float64) (int, string) {
	n, v := 1.0, 1.0
	for math.Abs(math.Pi-4*v) > epsilon {
		v += math.Pow(-1, n) / (2*n + 1)
		n += 1
	}
	return int(n), fmt.Sprintf("%.10f", 4*v)
}
