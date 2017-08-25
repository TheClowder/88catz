Code.load_file("24-aug-hellouer.exs", __DIR__)

ExUnit.start
ExUnit.configure trace: true, exclude: :pending

defmodule TestHellouer do
  use ExUnit.Case
  import Hellouer, only: [say_hello: 3]

  test "say hello to John Smith" do
    assert say_hello(["John", "Smith"], "Phoenix", "Arizona") == "Hello, John Smith! Welcome to Phoenix, Arizona!"
  end
end
