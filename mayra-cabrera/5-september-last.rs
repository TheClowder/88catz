// https://www.codewars.com/kata/last/train/rust
//
fn last<T: Clone>(slice: &[T]) -> T {
	return slice.last().unwrap().clone()
}
