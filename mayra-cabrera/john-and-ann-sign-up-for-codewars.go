//https://www.codewars.com/kata/john-and-ann-sign-up-for-codewars/train/go

package kata

func johnAndAnn(n int) ([]int, []int) {
	john := make([]int, n)
	john[0] = 0
	ann := make([]int, n)
	ann[0] = 1
	for i := 1; i < n; i++ {
		john[i] = i - ann[john[i-1]]
		ann[i] = i - john[ann[i-1]]
	}
	return john, ann
}

func Ann(n int) []int {
	_, a := johnAndAnn(n)
	return a
}
func John(n int) []int {
	j, _ := johnAndAnn(n)
	return j
}
func SumJohn(n int) int {
	j, _ := johnAndAnn(n)
	var sum = 0
	for i := 0; i < n; i++ {
		sum += j[i]
	}
	return sum
}
func SumAnn(n int) int {
	_, a := johnAndAnn(n)
	var sum = 0
	for i := 0; i < n; i++ {
		sum += a[i]
	}
	return sum
}
