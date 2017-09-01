extern crate nextitem;

#[test]
fn returns_expected() {
  assert_eq!(nextitem::next_item(&["Joe", "Bob", "Sally"], "Bob"), Some("Sally"));
  assert_eq!(nextitem::next_item(&[0, 1], 0), Some(1));
  assert_eq!(nextitem::next_item(&[0, 1], 1), None);
  assert_eq!(nextitem::next_item((1..10).collect::<Vec<_>>().as_slice(), 7), Some(8));  
}
