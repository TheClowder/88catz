# https://www.codewars.com/kata/whats-up-next/elixir

defmodule NextBigThing do
  def next_item(list, item) do
    list
    |> Enum.drop_while(fn x -> x != item end)
    |> Enum.drop(1)
    |> List.first
  end
end
