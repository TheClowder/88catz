extern crate jenny;

#[test]
fn greets_some_people_normally() {
  assert_eq!(jenny::greet("Jim"),   "Hello, Jim!");
  assert_eq!(jenny::greet("Jane"),  "Hello, Jane!");
  assert_eq!(jenny::greet("Simon"), "Hello, Simon!");
}

#[test]
fn greets_johnny_special() {
  assert_eq!(jenny::greet("Johnny"), "Hello, my love!");
}
