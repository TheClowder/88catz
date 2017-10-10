# https://www.codewars.com/kata/printer-errors

defmodule Printererror do
  
  def printer_error(s) do
    errors = Regex.scan(~r/[n-z]/, s)
    "#{Enum.count(errors)}/#{String.length(s)}"
  end
  
end
