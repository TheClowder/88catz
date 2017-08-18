// https://www.codewars.com/kata/all-unique/go

package kata

import (
	"strings"
)

func HasUniqueChar(str string) bool {
	array := strings.Split(str, "")
	uniqueArray := RemoveDuplicates(array)
	return len(array) == len(uniqueArray)
}

func RemoveDuplicates(elements []string) []string {
	duplicated := map[string]bool{}
	result := []string{}

	for v := range elements {
		if duplicated[elements[v]] != true {
			duplicated[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}
