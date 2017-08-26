# https://www.codewars.com/kata/559ac78160f0be07c200005a

defmodule Messy do
  def name_shuffler(str) do
    str 
    |> String.split(" ") 
    |> Enum.reverse() 
    |> Enum.join(" ")
  end
end
