# https://www.codewars.com/kata/count-by-x/train/elixir
defmodule Count do
  def count_by(x, n) do
    1..n |> Enum.map(fn(num) -> num * x end)
  end
end
