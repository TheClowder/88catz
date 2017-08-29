extern crate stringrepeat;

#[test]
fn example_tests() {
  assert_eq!(stringrepeat::repeat_str("a", 4), "aaaa");
  assert_eq!(stringrepeat::repeat_str("hello ", 3), "hello hello hello ");
  assert_eq!(stringrepeat::repeat_str("abc", 2), "abcabc");
}
