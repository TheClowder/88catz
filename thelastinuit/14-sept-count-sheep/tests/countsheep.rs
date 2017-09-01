extern crate countsheep;

#[test]
fn returns_correct_sheep_count() {
  assert_eq!(countsheep::count_sheep(&[false]), 0);
  assert_eq!(countsheep::count_sheep(&[true]), 1);
  assert_eq!(countsheep::count_sheep(&[true, false]), 1);
}
