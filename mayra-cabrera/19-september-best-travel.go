//https://www.codewars.com/kata/best-travel/go

package kata

func ChooseBestSum(t, k int, ls []int) int {
	n := len(ls)
	var dfs func(i int, ans int, m int) int

	dfs = func(i int, ans int, m int) int {
		if m == k {
			if ans > t {
				return -1
			}
			return ans
		}

		if i == n {
			return -1
		}
		res := -1

		for j := i; j < n; j++ {
			tmp := dfs(j+1, ans+ls[j], m+1)
			if tmp > res {
				res = tmp
			}
		}

		return res
	}

	return dfs(0, 0, 0)
}
