# https://www.codewars.com/kata/56f6ad906b88de513f000d96

defmodule Codewars.Perks do
  defstruct bonus: false

  defmacro is_hard_worker(bonus) do
    quote do: unquote(bonus) == true
  end

  defmacro is_slacker(bonus) do
    quote do: unquote(bonus) == false
  end
end

defmodule Codewars.Reward do
  import Codewars.Perks

  def bonus_time(salary, bonus) when is_hard_worker(bonus), do: "$#{salary * 10}"
  def bonus_time(salary, bonus) when is_slacker(bonus),     do: "$#{salary}"
end
