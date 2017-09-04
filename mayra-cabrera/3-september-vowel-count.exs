# https://www.codewars.com/kata/vowel-count/train/elixir

defmodule VowelCount do
  def get_count(str) do
    vowels = ["a", "e", "i", "o", "u"]
  	str
    |> String.graphemes
    |> Enum.count(&(&1 in vowels))
  end
end
