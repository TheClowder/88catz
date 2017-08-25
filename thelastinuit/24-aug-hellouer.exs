# https://www.codewars.com/kata/5302d846be2a9189af0001e4

defmodule Hellouer do
  def say_hello(name, city, state) do
    "Hello, #{name |> Enum.join(" ")}! Welcome to #{city}, #{state}!"
  end
end
