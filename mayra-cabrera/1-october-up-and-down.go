//https://www.codewars.com/kata/up-and-down

package kata

import "strings"

func Arrange(s string) string {
	arr := strings.Split(s, " ")

	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if (j+1)%2 == 0 {
				if len(arr[j]) < len(arr[j-1]) {
					arr[j], arr[j-1] = arr[j-1], arr[j]
				}
			} else {
				if len(arr[j]) > len(arr[j-1]) {
					arr[j], arr[j-1] = arr[j-1], arr[j]
				}
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		if (i+1)%2 == 0 {
			arr[i] = strings.ToUpper(arr[i])
		} else {
			arr[i] = strings.ToLower(arr[i])
		}
	}

	return strings.Join(arr, " ")
}
