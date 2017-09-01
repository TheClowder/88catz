extern crate ifcond;

#[test]
fn should_support_true() {
  assert_eq!(ifcond::_if(true, || 1, || 2), 1);
}

#[test]
fn should_support_false() {
  assert_eq!(ifcond::_if(false, || 1, || 2), 2);
}

#[test]
fn should_support_false_other_type() {
  assert_eq!(ifcond::_if(false, || 'a', || 'b'), 'b');
  assert_eq!(ifcond::_if(true, || "True", || "False"), "True");
}

#[test]
fn should_support_closures() {
  let mut true_was_set = false;
  let result = ifcond::_if(true, || true_was_set = true, || panic!("Fail"));
  assert!(true_was_set);
  assert_eq!(result, ())
}
