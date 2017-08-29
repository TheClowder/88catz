extern crate evenodd;

#[test]
fn returns_expected() {
  assert_eq!(evenodd::even_or_odd(0), "Even");
  assert_eq!(evenodd::even_or_odd(2), "Even");
  assert_eq!(evenodd::even_or_odd(1), "Odd");
  assert_eq!(evenodd::even_or_odd(7), "Odd");
}
