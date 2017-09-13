// https://www.codewars.com/kata/bouncing-balls/go

package kata

func BouncingBall(h, bounce, window float64) int {
	if h <= 0 || bounce <= 0 || bounce >= 1 || window >= h {
		return -1
	}
	result := 1

	for i := h * bounce; i > window; i *= bounce {
		result += 2
	}
	return result
}
