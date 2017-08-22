// https://www.codewars.com/kata/57f6ad55cca6e045d2000627/solutions/go
package kata

import "math"

func SquareOrSquareRoot(arr []int) []int {
	transformedNumbers := make([]int, len(arr))
	for position, number := range arr {
		squareRoot := int(math.Sqrt(float64(number)))
		if (squareRoot * squareRoot) == number {
			transformedNumbers[position] = squareRoot
		} else {
			transformedNumbers[position] = number * number
		}
	}
	return transformedNumbers
}
