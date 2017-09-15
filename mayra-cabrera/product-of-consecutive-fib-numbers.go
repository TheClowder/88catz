// https://www.codewars.com/kata/product-of-consecutive-fib-numbers/

package kata

func ProductFib(prod uint64) [3]uint64 {
	f1 := uint64(0)
	f2 := uint64(1)
	for f1*f2 < prod {
		f1, f2 = f2, f1+f2
	}
	success := uint64(0)
	if prod == f1*f2 {
		success = 1
	}
	return [3]uint64{f1, f2, success}
}
