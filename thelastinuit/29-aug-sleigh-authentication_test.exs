Code.load_file("29-aug-sleigh-authentication.exs", __DIR__)

ExUnit.start
ExUnit.configure trace: true, exclude: :pending

defmodule TestSleighAuthentication do
  use ExUnit.Case
  import SleighAuthentication, only: [authenticate?: 2]

  test "authenticates with correct credentials" do
    assert authenticate?("Santa Claus", "Ho Ho Ho!")
  end

  test "doesn't authenticate with incorrect credentials" do
    assert authenticate?("Santa", "Ho Ho Ho!") == false
    assert authenticate?("Santa Claus", "Ho Ho!") == false
    assert authenticate?("jhoffner", "CodeWars") == false
  end
end
