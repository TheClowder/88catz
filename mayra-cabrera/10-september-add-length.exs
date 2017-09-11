# https://www.codewars.com/kata/add-length/train/elixir

defmodule Marker do
  def add_length(str) do
    ls = String.split(str)
    Enum.map(ls, fn(x) -> "#{x} #{String.length(x)}" end)
  end
 end
