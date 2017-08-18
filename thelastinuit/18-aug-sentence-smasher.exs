defmodule SentenceSmasher do
  def smash(words) do
    Enum.join(words, " ")
  end
end
