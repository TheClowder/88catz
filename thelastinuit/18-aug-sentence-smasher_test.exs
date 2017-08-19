Code.load_file("18-aug-sentence-smasher.exs", __DIR__)

ExUnit.start
ExUnit.configure trace: true, exclude: :pending

defmodule TestSentenceSmasher do
  use ExUnit.Case
  import SentenceSmasher, only: [smash: 1]

  test "simple smashing" do
    assert smash(["hello"]) == "hello"
    assert smash(["hello", "world"]) == "hello world"
  end
end
