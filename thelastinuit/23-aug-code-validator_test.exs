Code.load_file("23-aug-code-validator.exs", __DIR__)

ExUnit.start
ExUnit.configure trace: true, exclude: :pending

defmodule TestCodeValidator do
  use ExUnit.Case
  import CodeValidator, only: [valid?: 1]

  test "basic cases" do
    assert valid?(123) == true
    assert valid?(248) == true
    assert valid?(8) == false
    assert valid?(321) == true
    assert valid?(9453) == false    
  end
end
