//https://www.codewars.com/kata/tortoise-racing

package kata

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}

	t := int(float64(g) / float64(v2-v1) * 3600.0)

	h := t / 3600
	t = t % 3600

	m := t / 60
	s := t % 60

	return [3]int{h, m, s}
}
