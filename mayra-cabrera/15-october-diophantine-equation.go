//https://www.codewars.com/kata/diophantine-equation

package kata

func Solequa(n int) [][]int {
	res := make([][]int, 0)
	var x, y float64
	diff := 1
	for x >= 0 && y >= 0 {
		x = float64(n+diff*diff) / float64(2*diff)
		y = (x - float64(diff)) / 2.0
		if float64(int(x)) == x && float64(int(y)) == y && (int(x)-2*int(y))*(int(x)+2*int(y)) == n && x >= 0 && y >= 0 {
			res = append(res, []int{int(x), int(y)})
		}
		diff++
	}
	return res
}
