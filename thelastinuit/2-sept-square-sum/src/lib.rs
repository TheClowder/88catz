// https://www.codewars.com/kata/515e271a311df0350d00000f

fn square(n: &i32) -> i32 { n * n }
fn accum(a: i32, x: i32) -> i32 { a + x }

pub fn square_sum(vec: Vec<i32>) -> i32 {
  vec.iter().map(square).fold(0, accum)
}
