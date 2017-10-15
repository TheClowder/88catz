//https://www.codewars.com/kata/esolang-interpreters-number-2-custom-smallfuck-interpreter

package kata

func pop(s []int) ([]int, int) {
	return s[:len(s)-1], s[len(s)-1]
}

func goToLoopEnd(code string, i int) int {
	openCount := 0
	for i < len(code) {
		switch code[i] {
		case '[':
			openCount++
		case ']':
			if openCount == 0 {
				return i
			} else {
				openCount--
			}
		}
		i++
	}
	return i
}

func Interpreter(code, tape string) string {
	mem := []rune(tape)
	i := 0
	m := 0
	var loopStack []int
	for i >= 0 && i < len(code) && m >= 0 && m < len(tape) {
		switch code[i] {
		case '<':
			m--
		case '>':
			m++
		case '*':
			if mem[m] == '0' {
				mem[m] = '1'
			} else {
				mem[m] = '0'
			}
		case '[':
			if mem[m] == '0' {
				i = goToLoopEnd(code, i+1)
			} else {
				loopStack = append(loopStack, i)
			}
		case ']':
			if mem[m] == '0' {
				loopStack, _ = pop(loopStack)
			} else {
				i = loopStack[0]
			}
		}
		i++
	}
	return string(mem)
}
