require "spec"
require "../src/*"

describe("Basic tests") do
  it "Should work for basic tests" do
    bonus_time(10000, true).should eq "$100000"
    bonus_time(25000, true).should eq "$250000"
    bonus_time(10000, false).should eq "$10000"
    bonus_time(60000, false).should eq "$60000"
    bonus_time(2, true).should eq "$20"
    bonus_time(78, false).should eq "$78"
    bonus_time(67890, true).should eq "$678900"
  end
end
