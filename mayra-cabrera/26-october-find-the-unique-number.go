//https://www.codewars.com/kata/find-the-unique-number-1

package kata

func FindUniq(arr []float32) float32 {
	size := len(arr)

	for i, number := range arr {
		right := arr[(i+1)%size]

		var left float32
		if i == 0 {
			left = arr[size-1]
		} else {
			left = arr[i-1]
		}

		if number != left && number != right {
			return number
		}
	}

	return arr[0]
}
