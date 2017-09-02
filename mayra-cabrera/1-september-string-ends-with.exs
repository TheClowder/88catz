# https://www.codewars.com/kata/string-ends-with/train/elixir

defmodule EndsWith do
  def solution(str, ending) do
    String.ends_with? str, ending
  end
end
