// https://www.codewars.com/kata/53da3dbb4a5168369a0000fe

pub fn even_or_odd(i: i32) -> &'static str {
  match i % 2 {
    0 => "Even",
    _ => "Odd",
  }
}

