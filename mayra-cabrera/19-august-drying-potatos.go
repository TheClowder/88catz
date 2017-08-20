https://www.codewars.com/kata/drying-potatoes/train/go

package kata

func Potatoes(p0, w0, p1 int) int {
	return w0 * (100 - p0) / (100 - p1)
}
