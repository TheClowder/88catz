//https://www.codewars.com/kata/esolang-interpreters-number-3-custom-paintf-star-star-k-interpreter
package kata

import (
	"errors"
	"fmt"
	"strings"
)

type Grid struct {
	Array         [][]int
	Width, Height int
	X, Y          int
}

func (g Grid) ToString() string {
	res := ""
	for _, r := range g.Array {
		for _, c := range r {
			res += fmt.Sprintf("%d", c)
		}
		res += "\r\n"
	}
	return strings.TrimSpace(res)
}

func (g *Grid) CurBit() int {
	return g.Array[g.Y][g.X]
}

func NewGrid(width, height int) Grid {
	g := Grid{Width: width, Height: height, Array: [][]int{}}
	for y := 0; y < height; y++ {
		g.Array = append(g.Array, make([]int, width))
	}
	return g
}

func (g *Grid) do(c string) error {
	switch c {
	case "n":
		g.Y -= 1
	case "s":
		g.Y += 1
	case "w":
		g.X -= 1
	case "e":
		g.X += 1
	case "*":
		if g.Array[g.Y][g.X] == 1 {
			g.Array[g.Y][g.X] = 0
		} else {
			g.Array[g.Y][g.X] = 1
		}
	case "[", "]":
		return nil
	default:
		return InvalidCommandErr
	}
	if g.Y < 0 {
		g.Y = g.Height - 1
	}
	if g.X < 0 {
		g.X = g.Width - 1
	}
	if g.Y >= g.Height {
		g.Y = 0
	}
	if g.X >= g.Width {
		g.X = 0
	}
	return nil
}

func Interpreter(code string, iterations, width, height int) string {
	g := NewGrid(width, height)
	pos := 0
	i := 0
	for i < iterations && pos < len(code) {
		if code[pos] == '[' && g.CurBit() == 0 || code[pos] == ']' && g.CurBit() == 1 {
			pos = findMatching(code[pos], pos, code) + 1
			i += 1
			continue
		}
		err := g.do(string(code[pos]))
		if err == nil {
			i += 1
		}
		pos += 1
	}
	return g.ToString()
}

func findMatching(char byte, pos int, code string) int {
	opposite := byte(']')
	direction := 1
	if char == ']' {
		opposite = byte('[')
		direction = -1
	}
	c := 1
	for c > 0 && pos > 0 && pos <= len(code) {
		pos += direction
		if code[pos] == char {
			c += 1
		}
		if code[pos] == opposite {
			c -= 1
		}
	}
	return pos
}
