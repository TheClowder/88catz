require "spec"
require "../src/*"

def testing(actual, expected)
  actual.should eq expected
end

describe "Basic tests" do
  it "Should work for basic tests" do
    testing(maps([1, 2, 3]), [2, 4, 6])
    testing(maps([4, 1, 1, 1, 4]), [8, 2, 2, 2, 8])
    testing(maps([2, 2, 2, 2, 2, 2]), [4, 4, 4, 4, 4, 4])
  end
end
