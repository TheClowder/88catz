//https://www.codewars.com/kata/piano-kata-part-1/

package kata

func BlackOrWhiteKey(keyPressCount int) string {
	octave := []int{0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1}
	keys := []string{"white", "black"}

	return keys[octave[(keyPressCount-1)%88%len(octave)]]
}
