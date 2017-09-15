require "spec"
require "../src/*"

class Player
  def initialize(@name : String)
  end

  def name
    @name
  end
end

describe "fixed tests" do
  it "should find the correct goose" do
    players = ("a".."z").to_a.map { |c| Player.new(c)}
    duck_duck_goose(players, 1).should eq "a"
    duck_duck_goose(players, 3).should eq "c"
    duck_duck_goose(players, 10).should eq "z"
    duck_duck_goose(players, 20).should eq "z"
    duck_duck_goose(players, 30).should eq "z"
    duck_duck_goose(players, 18).should eq "g"
    duck_duck_goose(players, 28).should eq "g"
    duck_duck_goose(players, 12).should eq "b"
    duck_duck_goose(players, 2).should eq "b"
    duck_duck_goose(players, 7).should eq "f"
  end
end
