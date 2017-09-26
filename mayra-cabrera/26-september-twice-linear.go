//https://www.codewars.com/kata/twice-linear/train/go

package main

func DblLinear(n int) int {
	u := []int{1}
	i1, i2 := 0, 0
	for len(u) <= n {
		c1, c2 := 2*u[i1]+1, 3*u[i2]+1
		if c1 < c2 {
			u = append(u, c1)
			i1++
		} else if c2 < c1 {
			u = append(u, c2)
			i2++
		} else {
			u = append(u, c1)
			i1++
			i2++
		}
	}
	return u[len(u)-1]
}
