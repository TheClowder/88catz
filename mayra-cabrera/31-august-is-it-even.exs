# https://www.codewars.com/kata/is-it-even/train/elixir

require Integer

defmodule Evenator do
 
  def even?(n) do
    Integer.is_even(round(n))
  end
end
