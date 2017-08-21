// https://www.codewars.com/kata/even-or-odd/train/go

package kata

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
