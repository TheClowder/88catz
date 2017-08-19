Code.load_file("17-aug-rock_paper_scissors.exs", __DIR__)

ExUnit.start
ExUnit.configure trace: true, exclude: :pending

defmodule TestRockPaperScissors do
  use ExUnit.Case
  import RockPaperScissors, only: [rps: 2]

  test "player 1 win" do
    assert rps("rock", "scissors") == "Player 1 won!"
    assert rps("scissors", "paper") == "Player 1 won!"
    assert rps("paper", "rock") == "Player 1 won!"
  end
  
  test "player 2 win" do
    assert rps("scissors", "rock") == "Player 2 won!"
    assert rps("paper", "scissors") == "Player 2 won!"
    assert rps("rock", "paper") == "Player 2 won!"
  end
  
  test "draw" do
    assert rps("rock", "rock") == "Draw!"
    assert rps("scissors", "scissors") == "Draw!"
    assert rps("paper", "paper") == "Draw!"
  end
end
