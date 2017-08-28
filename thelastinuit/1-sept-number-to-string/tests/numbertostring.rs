extern crate numbertostring;

#[test]
fn returns_number_as_a_string() {
  assert_eq!(numbertostring::number_to_string(67), "67");
  assert_eq!(numbertostring::number_to_string(1+2), "3");
}
