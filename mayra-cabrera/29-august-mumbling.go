package kata

import "strings"

func Accum(s string) string {
	var output string
	for i, r := range s {
		repetition := strings.Repeat(string(r), i+1)
		output = output + strings.ToLower(repetition) + "-"
	}
	output = strings.Title(output)
	return strings.TrimSuffix(output, "-")
}
