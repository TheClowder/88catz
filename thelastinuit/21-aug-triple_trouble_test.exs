Code.load_file("21-aug-triple_trouble.exs", __DIR__)

ExUnit.start
ExUnit.configure trace: true, exclude: :pending

defmodule Codewars.TestWeirdString do
  use ExUnit.Case
  import Codewars.WeirdString, only: [triple_trouble: 3]

  test "basic cases" do
    assert triple_trouble("this", "test", "lock") == "ttlheoiscstk"
    assert triple_trouble("aa","bb","cc") == "abcabc"
    assert triple_trouble("Bm", "aa", "tn") == "Batman"
    assert triple_trouble("LLh", "euo", "xtr") == "LexLuthor"
    assert triple_trouble("aaa", "bbb", "ccc") == "abcabcabc"
    assert triple_trouble("aaaaaa", "bbbbbb", "cccccc") == "abcabcabcabcabcabc"
    assert triple_trouble("burn",  "reds",  "roll") == "brrueordlnsl"
    assert triple_trouble("Sea", "urn", "pms") == "Supermans" 
    assert triple_trouble("LLh", "euo", "xtr") == "LexLuthor"
  end
end
