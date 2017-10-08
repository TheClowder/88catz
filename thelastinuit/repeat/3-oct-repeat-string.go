// https://www.codewars.com/kata/57a0e5c372292dd76d000d7e

package repeat

import (
  "strings"
)

func RepeatStr(repititions int, value string) string {
  return strings.Repeat(value, repititions)
}
