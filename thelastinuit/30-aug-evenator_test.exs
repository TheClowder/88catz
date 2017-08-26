Code.load_file("30-aug-evenator.exs", __DIR__)

ExUnit.start
ExUnit.configure trace: true, exclude: :pending

defmodule TestSolution do
  use ExUnit.Case
  import Evenator, only: [even?: 1]

  test "basic cases" do
    assert even?(0) == true
    assert even?(0.5) == false
    assert even?(1) == false
    assert even?(2) == true
    assert even?(-4) == true
  end
end
