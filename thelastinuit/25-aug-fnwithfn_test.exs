Code.load_file("25-aug-fnwithfn.exs", __DIR__)

ExUnit.start
ExUnit.configure trace: true, exclude: :pending

defmodule TestFnwithfn do
  use ExUnit.Case
  import Fnwithfn, only: [always: 1]

  test "return a function which return 3" do
    assert always(3).() == 3
  end
end
