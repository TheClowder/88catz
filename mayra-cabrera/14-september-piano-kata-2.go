//https://www.codewars.com/kata/piano-kata-part-2/go

package kata

func WhichNote(key int) string {
	keys := []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	return keys[((key-1)%88)%len(keys)]
}
