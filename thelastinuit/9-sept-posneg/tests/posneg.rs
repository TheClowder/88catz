extern crate posneg;

#[test]
fn returns_expected() {
  let test_data1 = vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15];
  let expected1 = vec![10, -65];
  assert_eq!(posneg::count_positives_sum_negatives(test_data1), expected1);    
}
