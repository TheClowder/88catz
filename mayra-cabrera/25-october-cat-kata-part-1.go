//https://www.codewars.com/kata/cat-kata-part-1

package kata

import "math"

func PeacefulYard(yard []string, minDistance int) bool {
	var catsStock [][]int
	for row := 0; row < len(yard); row++ {
		for col := 0; col < len(yard[0]); col++ {
			if yard[row][col] != '-' {
				catsStock = append(catsStock, []int{row, col})
			}
		}
	}
	for i := 0; i < len(catsStock)-1; i++ {
		for j := i + 1; j < len(catsStock); j++ {
			if math.Sqrt(math.Pow(float64(catsStock[i][0]-catsStock[j][0]), float64(2))+math.Pow(float64(catsStock[i][1]-catsStock[j][1]), float64(2))) < float64(minDistance) {
				return false
			}
		}
	}
	return true
}
