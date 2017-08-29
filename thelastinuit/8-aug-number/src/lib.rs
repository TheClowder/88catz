// https://www.codewars.com/kata/5648b12ce68d9daa6b000099

fn accum(acc: i32, flux: &(i32, i32)) -> i32 {
  acc + flux.0 - flux.1
}

fn number(bus_stops: &[(i32,i32)]) -> i32 {
  bus_stops.iter().fold(0, accum)
}
