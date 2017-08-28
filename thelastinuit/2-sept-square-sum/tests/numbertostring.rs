extern crate squaresum;

#[test]
fn returns_expected() {
  assert_eq!(squaresum::square_sum(vec![1, 2]), 5);
  assert_eq!(squaresum::square_sum(vec![-1, -2]), 5);
  assert_eq!(squaresum::square_sum(vec![5, 3, 4]), 50);
}
