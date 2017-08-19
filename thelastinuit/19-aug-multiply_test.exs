Code.load_file("19-aug-multiply.exs", __DIR__)

ExUnit.start
ExUnit.configure trace: true, exclude: :pending

defmodule TestMultiply do
  use ExUnit.Case
  import Multiply, only: [multiply: 2]

  test "multiply two numbers" do
    assert multiply(1,2) == 2
    assert multiply(3,8) == 24
  end
end
