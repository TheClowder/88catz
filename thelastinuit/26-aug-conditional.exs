# https://www.codewars.com/kata/54147087d5c2ebe4f1000805

defmodule Conditional do
  def _if(bool, fthen, felse) do
    if bool, do: fthen.(), else: felse.()
  end
end
