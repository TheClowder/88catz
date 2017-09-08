defmodule Inverter do
  def invert(list) do
    Enum.map(list, fn(x) -> x - 1 end)
  end
end
