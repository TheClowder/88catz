//https://www.codewars.com/kata/simple-fun-number-122-string-constructing

package kata

func StringConstructing(a, s string) int {
	word := a
	step := 1
	for i, c := range s {
		if len(word) < len(s) {
			word += a
			step++
		}

		for word[i] != byte(c) {
			word = word[:i] + word[i+1:]
			step++
			if len(word) < len(s) {
				word += a
				step++
			}
		}
	}
	return step + len(word) - len(s)
}
