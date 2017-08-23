# https://www.codewars.com/kata/56a25ba95df27b7743000016

defmodule CodeValidator do
  def valid?(code) do
    match = Integer.to_string(code, 10) |> (&Regex.scan(~r/^1|^2|^3/, &1)).() |> List.first

    if match, do: true, else: false
  end
end
