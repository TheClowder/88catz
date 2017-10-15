//https://www.codewars.com/kata/playing-with-digits/

package kata

import "math"

func DigPow(n, p int) int {
	total := 0
	digit := n
	l := int(math.Log10(float64(n))) + p
	for i := l; n > 0; i-- {
		total += int(math.Pow(float64(n%10), float64(i)))
		n /= 10
	}
	if total%digit == 0 {
		return total / digit
	} else {
		return -1
	}
}
