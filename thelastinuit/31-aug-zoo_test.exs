Code.load_file("31-aug-zoo.exs", __DIR__)

ExUnit.start
ExUnit.configure trace: true, exclude: :pending

defmodule Codewars.TestZoo do
  use ExUnit.Case
  import Codewars.Zoo, only: [fix_the_meerkat: 1]

  test "basic cases" do
    assert fix_the_meerkat({"tail", "body", "head"}) == {"head", "body", "tail"}
    assert fix_the_meerkat({"tails", "body", "heads"}) == {"heads", "body", "tails"}
    assert fix_the_meerkat({"bottom", "middle", "top"}) == {"top", "middle", "bottom"}
    assert fix_the_meerkat({"lower legs", "torso", "upper legs"}) == {"upper legs", "torso", "lower legs"}
    assert fix_the_meerkat({"ground", "rainbow", "sky"}) == {"sky", "rainbow", "ground"}
  end
end
