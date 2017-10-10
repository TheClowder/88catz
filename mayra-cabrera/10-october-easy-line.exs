# https://www.codewars.com/kata/easy-line/train/elixir

defmodule Easyline do
  
  def easyline(n) do
    div(factorial(2*n), factorial(n) * factorial(n))
  end
  
  def factorial(0), do: 1
  def factorial(n), do: n * factorial(n-1)
end
