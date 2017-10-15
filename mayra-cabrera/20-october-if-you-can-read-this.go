//https://www.codewars.com/kata/if-you-can-read-this-dot-dot-dot

package kata

import "strings"

func ToNato(words string) string {
	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu"}
	words = strings.Replace(words, " ", "", -1)
	result := ""

	for index, value := range words {

		var natoWord string
		for _, innerValue := range nato {
			if strings.HasPrefix(innerValue, strings.ToUpper(string(value))) {
				natoWord = innerValue
				break
			}

			if natoWord == "" {
				natoWord = string(value)
			}
		}

		if index < len(words)-1 {
			result += natoWord + " "
		} else {
			result += natoWord
		}

	}

	return result
}
