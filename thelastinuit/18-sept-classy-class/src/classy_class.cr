# https://www.codewars.com/kata/55a144eff5124e546400005a
class Person
  def initialize(@name : String, @age : Int32)
  end

  def info
    "#{@name}s age is #{@age}"
  end
end