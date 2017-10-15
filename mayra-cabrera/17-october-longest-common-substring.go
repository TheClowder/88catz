//https://www.codewars.com/kata/longest-common-subsequence

package kata

func lcs(s1, s2 *string, i, j int) string {
	var cache = make(map[string]string)
	k := string(i) + string(j) + *s1 + *s2
	if _, ok := cache[k]; !ok {
		if i == 0 || j == 0 {
			cache[k] = ""
		} else if (*s1)[i-1] == (*s2)[j-1] {
			cache[k] = lcs(s1, s2, i-1, j-1) + string((*s1)[i-1])
		} else {
			ss1 := lcs(s1, s2, i, j-1)
			ss2 := lcs(s1, s2, i-1, j)
			if len(ss1) > len(ss2) {
				cache[k] = ss1
			} else {
				cache[k] = ss2
			}
		}
	}
	return cache[k]
}

func LCS(s1, s2 string) string {
	return lcs(&s1, &s2, len(s1), len(s2))
}
