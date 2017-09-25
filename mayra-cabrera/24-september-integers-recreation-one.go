//http://www.codewars.com/kata/integers-recreation-one
package kata

import "math"

func ListSquared(m, n int) [][]int {
	res := make([][]int, 0)
	for i := m; i <= n; i++ {
		s := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				s += j * j
			}
		}
		if math.Mod(math.Sqrt(float64(s)), 1.) == 0 {
			res = append(res, []int{i, s})
		}
	}
	return res
}
