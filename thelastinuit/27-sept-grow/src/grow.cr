# https://www.codewars.com/kata/57f780909f7e8e3183000078

def grow(x) 
  x.reduce(1) { |acc, i| acc * i }  
end
