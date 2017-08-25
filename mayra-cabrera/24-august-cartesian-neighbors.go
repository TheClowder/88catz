// https://www.codewars.com/kata/cartesian-neighbors/train/go

package kata

func CartesianNeighbor(x, y int) [][]int {
	var neighbors [][]int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			neighbors = append(neighbors, []int{x + i, y + j})
		}
	}
	return neighbors
}
