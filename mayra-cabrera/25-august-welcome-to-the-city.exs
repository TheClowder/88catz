# http://www.codewars.com/kata/welcome-to-the-city/elixir

defmodule Hellouer do
  def say_hello(name, city, state) do
    complete_name = Enum.join(name, " ")
    "Hello, #{complete_name}! Welcome to #{city}, #{state}!"
  end
end
