// https://www.codewars.com/kata/57a0e5c372292dd76d000d7e 
use std::iter::repeat;

pub fn repeat_str(src: &str, count: usize) -> String {
  repeat(src).take(count).collect()
}
