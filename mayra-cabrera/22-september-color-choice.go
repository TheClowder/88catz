//https://www.codewars.com/kata/color-choice/train/go

package kata

import "math/big"

func CheckChoose(m, n int) int {
	if n == 0 {
		return 0
	}
	for x := 0; x < n; x++ {
		zBigInt := big.NewInt(0)
		zBigInt.Binomial(int64(n), int64(x))
		v := int(zBigInt.Int64())
		if v == m {
			return x
		}
	}
	return -1
}
