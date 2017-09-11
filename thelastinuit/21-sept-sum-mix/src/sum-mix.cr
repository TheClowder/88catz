# https://www.codewars.com/kata/57eaeb9578748ff92a000009

def sum_mix(x)
  x.reduce(0) { |acc, i| acc + i.to_i }
end
