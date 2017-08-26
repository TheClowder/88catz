# https://www.codewars.com/kata/52adc142b2651f25a8000643

defmodule SleighAuthentication do
  def authenticate?(name, password) do
    if name === "Santa Claus" and password === "Ho Ho Ho!", do: true, else: false
  end
end
