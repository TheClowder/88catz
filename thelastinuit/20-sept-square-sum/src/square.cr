# https://www.codewars.com/kata/515e271a311df0350d00000f

def square_sum(arr)
  arr.reduce(0) { |acc, i| acc + i * i }
end
