# https://www.codewars.com/kata/how-green-is-my-valley

defmodule Makevalley do

  def make_valley(a) do
    sorted  = Enum.sort(a, &(&1 > &2))
    {_, tl} = Enum.split(sorted,1)
    left    = Enum.take_every(sorted, 2)
    right   = Enum.take_every(tl, 2) |> Enum.reverse
    left ++ right
  end
  
end
