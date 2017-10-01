//https://www.codewars.com/kata/gap-in-primes/go

package kata

import "math"

func Gap(g, m, n int) []int {
	var primes []int
	for i := m; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	previous := -1
	for _, nr := range primes {
		if previous == -1 {
			previous = nr
			continue
		}
		if nr-previous == g {
			return []int{previous, nr}
		}
		previous = nr
	}
	return nil
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
