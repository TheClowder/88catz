# https://www.codewars.com/kata/going-to-the-cinema

defmodule Movie do
  def movie(card, ticket, perc, count \\ 1) do
    nCard = card + (ticket * :math.pow(perc, count))

    if Float.ceil(nCard / 1) < (ticket * count) do
        count    
    else
        movie(nCard, ticket, perc, count + 1)    
    end
  end
end
