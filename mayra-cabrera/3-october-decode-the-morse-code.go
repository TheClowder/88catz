//https://www.codewars.com/kata/decode-the-morse-code

package kata

import "strings"

func DecodeMorse(morseCode string) string {
	sentence := strings.Split(strings.Trim(morseCode, " "), "  ")
	for i, character := range sentence {
		word := strings.Split(character, " ")
		for j, c := range word {
			word[j] = MORSE_CODE[c]
		}
		sentence[i] = strings.Join(word, "")
	}
	return strings.Join(sentence, " ")
}
