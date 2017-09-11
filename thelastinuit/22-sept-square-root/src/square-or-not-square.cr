# https://www.codewars.com/kata/57f6ad55cca6e045d2000627

def square_or_square_root(arr)
  arr.map { |i| 
    if Math.sqrt(i) % 1 == 0
      Math.sqrt(i).to_i
    else
      i * i
    end
  }
end
