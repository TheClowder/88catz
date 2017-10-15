//https://www.codewars.com/kata/n-point-crossover/

package kata

import "sort"

func Crossover(ns []int, xs []int, ys []int) ([]int, []int) {
	if len(ns) == 0 {
		return xs, ys
	}
	sort.Ints(ns)
	prevC := -1
	for _, i := range ns {
		if i == prevC {
			continue
		}
		prevC = i
		temporal := make([]int, len(xs[i:]))
		copy(temporal, xs[i:])
		copy(xs[i:], ys[i:])
		copy(ys[i:], temporal)
	}
	return xs, ys
}
