extern crate alternatingcase;

#[test]
fn example_tests() {
  assert_eq!("HELLO WORLD", alternatingcase::to_alternating_case("hello world"));
  assert_eq!("hello world", alternatingcase::to_alternating_case("HELLO WORLD"));
  assert_eq!("HELLO world", alternatingcase::to_alternating_case("hello WORLD"));
  assert_eq!("hEllO wOrld", alternatingcase::to_alternating_case("HeLLo WoRLD"));
  assert_eq!("Hello World", alternatingcase::to_alternating_case(&alternatingcase::to_alternating_case("Hello World")[..]));
  assert_eq!("12345", alternatingcase::to_alternating_case("12345"));
  assert_eq!("1A2B3C4D5E", alternatingcase::to_alternating_case("1a2b3c4d5e"));
  assert_eq!("sTRING.tOaLTERNATINGcASE", alternatingcase::to_alternating_case("String.ToAlternatingCase"));
}
