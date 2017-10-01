//https://www.codewars.com/kata/consecutive-strings
package kata

func LongestConsec(strarr []string, k int) string {
	longest := ""
	longestLength := 0

	for i := 0; i < len(strarr)-k+1; i++ {
		var word string
		for j := 0; j < k; j++ {
			word += strarr[i+j]
		}

		wordLength := len(word)
		if wordLength > longestLength {
			longestLength = wordLength
			longest = word
		}
	}
	return longest
}
