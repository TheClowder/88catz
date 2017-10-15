//https://www.codewars.com/kata/number-of-integer-partitions

package kata

func partitions(n int, cache map[int]int) int {
	if n <= 1 {
		return 1
	}
	if _, value := cache[n]; !value {
		m := n - 1
		size := 0
		step := 2
		for m >= 0 {
			prevSize := partitions(m, cache)
			if step/2%2 == 1 {
				size += prevSize
			} else {
				size -= prevSize
			}
			if step%2 == 1 {
				m -= step
			} else {
				m -= step / 2
			}
			step++
		}
		cache[n] = size
	}
	return cache[n]
}

func Partitions(n int) int {
	cache := make(map[int]int, n-1)
	return partitions(n, cache)
}
