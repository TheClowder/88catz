//https://www.codewars.com/kata/fizz-buzz-cuckoo-clock/go

package kata

import "strings"
import "strconv"

func FizzBuzzCuckooClock(time string) string {
	sp := strings.Split(time, ":")
	hour, _ := strconv.Atoi(sp[0])
	min, _ := strconv.Atoi(sp[1])
	if min == 0 {
		n := hour % 12
		if n == 0 {
			n = 12
		}
		return strings.TrimSpace(strings.Repeat("Cuckoo ", n))
	}
	if min%30 == 0 {
		return "Cuckoo"
	}
	if min%5 == 0 && min%3 == 0 {
		return "Fizz Buzz"
	}
	if min%5 == 0 {
		return "Buzz"
	}
	if min%3 == 0 {
		return "Fizz"
	}
	return "tick"
}
