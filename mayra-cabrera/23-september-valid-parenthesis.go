//https://www.codewars.com/kata/valid-parentheses

package main

import "fmt"

func ValidParentheses(parens string) bool {
	p := 0
	for _, c := range parens {
		if p < 0 {
			return false
		}
		if c == '(' {
			p++
		} else {
			p--
		}
	}
	return p == 0
}
