// 

pub fn last<T: Clone>(slice: &[T]) -> T {
  slice.last().unwrap().clone()
}
