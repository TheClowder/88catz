# https://www.codewars.com/kata/557af9418895e44de7000053

defmodule Repeater do
  def repeat_it(str, n) when is_bitstring(str), do: String.duplicate(str, n)
  def repeat_it(_str, _n), do: "Not a string"
end
