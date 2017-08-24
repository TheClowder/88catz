// https://www.codewars.com/kata/511f11d355fe575d2c000001/train/go

package kata

import "sort"

func TwoOldestAges(ages []int) [2]int {
	sort.Slice(ages, func(i, j int) bool {
		return ages[i] > ages[j]
	})
	return [2]int{ages[1], ages[0]}
}
