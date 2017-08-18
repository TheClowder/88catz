# https://www.codewars.com/kata/53dc23c68a0c93699800041d

defmodule SentenceSmasher do
  def smash(words) do
    Enum.join(words, " ")
  end
end
