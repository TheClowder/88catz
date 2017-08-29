// https://www.codewars.com/kata/56efc695740d30f963000557

fn mutate(c: char) -> String {
  if c.to_string().find(char::is_uppercase) == Some(0) {
    c.to_string().to_lowercase()
  } else {
    c.to_string().to_uppercase()
  }
}

pub fn to_alternating_case(s: &str) -> String {
  s.chars().map(mutate).collect()
}
