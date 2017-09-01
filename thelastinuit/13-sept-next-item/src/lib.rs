// https://www.codewars.com/kata/542ebbdb494db239f8000046

pub fn next_item<T: PartialEq<T> + Clone>(slice: &[T], find: T) -> Option<T> {
  slice.iter().skip_while(|&x| *x != find).nth(1).cloned()
}
