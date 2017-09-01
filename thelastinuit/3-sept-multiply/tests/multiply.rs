extern crate multiply;

#[test]
fn returns_expected() {
  assert_eq!(multiply::multiply(3, 5), 15)
}
