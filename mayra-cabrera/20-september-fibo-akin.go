//https://www.codewars.com/kata/fibo-akin/train/go

package kata

var h map[int]int = map[int]int{
	0: 1,
	1: 1,
}

func u(n int) int {
	if _, ok := h[n]; !ok {
		h[n] = h[n-h[n-1]] + h[n-h[n-2]]
	}
	return h[n]
}
func LengthSupUk(n, k int) int {
	count := 0
	for i := 0; i < n; i++ {
		if u(i) >= k {
			count++
		}
	}
	return count
}
func Comp(n int) int {
	count := 0
	for i := 1; i < n; i++ {
		if u(i-1) > u(i) {
			count++
		}
	}
	return count
}
