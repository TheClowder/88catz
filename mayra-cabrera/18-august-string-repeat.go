// https://www.codewars.com/kata/string-repeat/train/go

package kata

import "strings"

func RepeatStr(repititions int, value string) string {
	return strings.Repeat(value, repititions)
}
