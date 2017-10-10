# https://www.codewars.com/kata/parts-of-a-list

defmodule Partlist do

  def part_list(a) do 
    for count <- 1..(length(a) - 1) do
      with {head, tail} <- Enum.split(a, count) do
        [Enum.join(head, " "), Enum.join(tail, " ")]
      end
    end
  end
  
end
