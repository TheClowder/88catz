# https://www.codewars.com/kata/welcome/train/elixir

defmodule Greeter do

	
  def greet(language) do
    Keyword.get(languages, String.to_atom(language), "Welcome")
  end
  
  def languages do
  [
    english: "Welcome",
    czech: "Vitejte",
    danish: "Velkomst",
    dutch: "Welkom",
    estonian: "Tere tulemast",
    finnish: "Tervetuloa",
    flemish: "Welgekomen",
    french: "Bienvenue",
    german: "Willkommen",
    irish: "Failte",
    italian: "Benvenuto",
    latvian: "Gaidits",
    lithuanian: "Laukiamas",
    polish: "Witamy",
    spanish: "Bienvenido",
    swedish: "Valkommen",
    welsh: "Croeso"
   ]
  end
end
