# https://www.codewars.com/kata/sum-of-two-lowest-positive-integers/train/elixir
defmodule SmallSummer do
  def sum_two_smallest_numbers(numbers) do
  	sort_numbers = Enum.sort(numbers)
    first = Enum.at(sort_numbers, 0)
    second = Enum.at(sort_numbers, 1)
    first + second
  end
end
