Code.load_file("28-aug-repeater.exs", __DIR__)

ExUnit.start
ExUnit.configure trace: true, exclude: :pending

defmodule TestSolution do
  use ExUnit.Case
  import Repeater, only: [repeat_it: 2]

  test "basic tests" do
    assert repeat_it("*", 3) == "***"
    assert repeat_it("Hello", 11) == "HelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHello"
  end
end
