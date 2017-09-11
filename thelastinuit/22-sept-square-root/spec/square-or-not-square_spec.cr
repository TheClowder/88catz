require "spec"
require "../src/*"

describe "Basic tests" do
  it "With #{[4, 3, 9, 7, 2, 1]} should return [2, 9, 3, 49, 4, 1]" { square_or_square_root([4, 3, 9, 7, 2, 1]).should eq [2, 9, 3, 49, 4, 1]} 
  it "With #{[100, 101, 5, 5, 1, 1]} should return [10, 10201, 25, 25, 1, 1]" { square_or_square_root([100, 101, 5, 5, 1, 1]).should eq [10, 10201, 25, 25, 1, 1]} 
  it "With #{[1, 2, 3, 4, 5, 6]} should return 41" { square_or_square_root([1, 2, 3, 4, 5, 6]).should eq [1, 4, 9, 2, 25, 36]} 
end
