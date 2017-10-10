# https://www.codewars.com/kata/speed-control

defmodule Speedcontrol do
  def gps(s, [a, b | rest]) do
    average = Float.floor((b - a) * 3600.0 / s)
    Enum.max([average, gps(s, [b] ++ rest)])
  end
  
  def gps(_, _) do
    0
  end
end
