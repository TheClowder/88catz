Code.load_file("27-aug-messy.exs", __DIR__)

ExUnit.start
ExUnit.configure trace: true, exclude: :pending

defmodule TestMessy do
  use ExUnit.Case
  import Messy, only: [name_shuffler: 1]

  test "basic tests" do
    assert name_shuffler("john McClane") == "McClane john"
    assert name_shuffler("Mary jeggins") == "jeggins Mary"
    assert name_shuffler("tom jerry") == "jerry tom"
  end
end
