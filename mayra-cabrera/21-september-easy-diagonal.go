//https://www.codewars.com/kata/easy-diagonal/

package kata

import "math/big"

func Diagonal(n, p int) int {
	var z big.Int
	return int(z.Binomial(int64(n+1), int64(p+1)).Int64())
}
