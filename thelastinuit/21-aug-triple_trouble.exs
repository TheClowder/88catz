# https://www.codewars.com/kata/5704aea738428f4d30000914

defmodule Codewars.WeirdString do
  def triple_trouble(one, two, three) do
    # For fancy chars, use graphemes. it's slow, though.
    chars_one = one |> String.codepoints
    chars_two = two |> String.codepoints
    chars_three = three |> String.codepoints
     
    Enum.map(0..length(chars_one) - 1, fn (index) ->  
      [
        chars_one |> Enum.at(index),
        chars_two |> Enum.at(index),
        chars_three |> Enum.at(index)
      ] |> Enum.join
    end) |> Enum.join
  end
end
