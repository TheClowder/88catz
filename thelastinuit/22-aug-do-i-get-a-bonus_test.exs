Code.load_file("22-aug-do-i-get-a-bonus.exs", __DIR__)

ExUnit.start
ExUnit.configure trace: true, exclude: :pending

defmodule Codewars.TestReward do
  use ExUnit.Case
  import Codewars.Reward, only: [bonus_time: 2]

  test "basic cases" do
    assert bonus_time(10000, true) == "$100000"
    assert bonus_time(25000, true) == "$250000"
    assert bonus_time(10000, false) == "$10000"
    assert bonus_time(60000, false) == "$60000"
    assert bonus_time(2, true) == "$20"
    assert bonus_time(78, false) == "$78"
    assert bonus_time(67890, true) == "$678900"
  end
end
