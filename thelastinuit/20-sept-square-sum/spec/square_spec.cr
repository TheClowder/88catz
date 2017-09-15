require "spec"
require "../src/*"

describe "Basic tests" do
  it "works!" do
    square_sum([1]).should eq 1
    square_sum([1,2]).should eq 5
    square_sum([0, 3, 4, 5]).should eq 50
    square_sum([0,-3,-4,-5]).should eq 50
    square_sum([] of Int32).should eq 0
  end
end
