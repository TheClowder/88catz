require "spec"
require "../src/*"

describe("Basic tests") do
  it "With Fox" { correct_tail("Fox", "x").should eq true }
  it "With Rhino" { correct_tail("Rhino", "o").should eq true }
  it "With Meerkat" { correct_tail("Meerkat", "t").should eq true }
  it "With Emu" { correct_tail("Emu", "t").should eq false }
  it "With Badger" { correct_tail("Badger", "s").should eq false }
  it "With Giraffe" { correct_tail("Giraffe", "d").should eq false }
end
