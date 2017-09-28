//https://www.codewars.com/kata/decode-the-morse-code-advanced/train/go

package main

import "strings"

func DecodeBits(bits string) string {
	i, unit, unit0 := 0, 0, 0
	bits = strings.Trim(bits, "0")
	for _, c := range bits + "0" {
		if c == '1' {
			i++
			if unit == 0 || unit0 > 0 && unit0 < unit {
				unit = unit0
			}
			unit0 = 0
		} else {
			unit0++
			if unit == 0 || i > 0 && i < unit {
				unit = i
			}
			i = 0
		}
	}
	bits = replace(bits, "0", "  ", unit*7)
	bits = replace(bits, "0", " ", unit*3)
	bits = replace(bits, "1", "-", unit*3)
	bits = replace(bits, "1", ".", unit)
	bits = replace(bits, "0", "", unit)
	return bits
}

func replace(s, b, c string, r int) string {
	return strings.Replace(s, strings.Repeat(b, r), c, -1)
}

func DecodeMorse(morseCode string) (result string) {
	s := strings.Split(morseCode, "  ")
	for i, w := range s {
		for _, c := range strings.Split(w, " ") {
			result += MORSE_CODE[c]
		}
		if i+1 != len(s) {
			result += " "
		}
	}
	return
}
