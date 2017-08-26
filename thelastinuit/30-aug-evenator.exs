# https://www.codewars.com/kata/555a67db74814aa4ee0001b5

require Integer

defmodule Evenator do
    def even?(n) when is_integer(n), do: Integer.is_even(n)
    def even?(n) when is_float(n), do: false
end
