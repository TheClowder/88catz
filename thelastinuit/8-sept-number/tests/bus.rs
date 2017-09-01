extern crate bus;

#[test]
fn returns_expected() {
  assert_eq!(bus::number(&[(10,0),(3,5),(5,8)]), 5);
  assert_eq!(bus::number(&[(3,0),(9,1),(4,10),(12,2),(6,1),(7,10)]), 17);
  assert_eq!(bus::number(&[(3,0),(9,1),(4,8),(12,2),(6,1),(7,8)]), 21);
}
