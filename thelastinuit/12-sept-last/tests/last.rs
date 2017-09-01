extern crate last;

#[test]
fn should_work_for_non_empty_list_of_integers() {
  assert_eq!(last::last(&[1, 2, 3]), 3);
}

#[test]
fn should_work_for_non_empty_list_of_strings() {
  assert_eq!(last::last(&["a", "b"]), "b");
}
