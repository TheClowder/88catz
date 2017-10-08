// https://www.codewars.com/kata/57f6ad55cca6e045d2000627

package square

import "math"

func SquareOrSquareRoot(arr []int) []int{
  new_arr := make([]int, len(arr))

  for index, value := range arr { 
	if int(math.Sqrt(float64(value))) * int(math.Sqrt(float64(value))) == value {
	  new_arr[index] = int(math.Sqrt(float64(value)))
	} else {
	  new_arr[index] = value * value
	}
  }

  return new_arr
}
