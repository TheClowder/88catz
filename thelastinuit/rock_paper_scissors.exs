defmodule RockPaperScissors do
  def rps(p1, p2) do
    outcomes = %{
      "paper"    => %{
        "rock"     => "Player 1 won!", 
        "scissors" => "Player 2 won!",
        "paper"    => "Draw!"
      },
      "rock"     => %{
        "scissors" => "Player 1 won!", 
        "paper"    => "Player 2 won!",
        "rock"     => "Draw!"
      },
      "scissors" => %{
        "paper"    => "Player 1 won!", 
        "rock"     => "Player 2 won!",
        "scissors" => "Draw!"
      },
    }

    outcomes[p1][p2]
  end
end
