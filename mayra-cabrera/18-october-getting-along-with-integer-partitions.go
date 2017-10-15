//https://www.codewars.com/kata/getting-along-with-integer-partitions/

package kata

import (
	"fmt"
	"sort"
)

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func removeDuplicates(a []int) []int {
	result := []int{}
	seen := map[int]int{}
	for _, value := range a {
		if _, ok := seen[value]; !ok {
			result = append(result, value)
			seen[value] = value
		}
	}
	return result
}

func average(a []int) float64 {
	var total float64 = 0
	for _, value := range a {
		total += float64(value)
	}
	return total / float64(len(a))
}

func partAux(s, k int) [][]int {
	k0 := min(s, k)
	res := [][]int{}
	n := k0
	var r int
	for ; n > 0; n-- {
		r = s - n
		if r > 0 {
			arr := partAux(r, n)
			for i := 0; i < len(arr); i++ {
				t := arr[i]
				t = append(t, n)
				res = append(res, t)
			}
		} else {
			res = append(res, []int{n})
		}
	}
	return res
}
func arrayProds(arr [][]int) []int {
	var prods []int
	for _, a := range arr {
		p := 1
		for _, val := range a {
			p *= val
		}
		prods = append(prods, p)
	}
	return prods
}
func Part(n int) string {
	prods := removeDuplicates(arrayProds(partAux(n, n)))
	sort.Ints(prods)
	avg := average(prods)
	lg := len(prods)
	rge := prods[lg-1] - prods[0]
	md := float64(prods[(lg-1)/2]+prods[lg/2]) / 2.0
	return fmt.Sprintf("Range: %d Average: %.2f Median: %.2f", rge, avg, md)
}
