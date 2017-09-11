require "spec"
require "../src/*"

describe "Basic grows" do
  grow([1, 2, 3]).should eq 6
  grow([4, 1, 1, 1, 4]).should eq 16
  grow([2, 2, 2, 2, 2, 2]).should eq 64
  grow([1]).should eq 1
  grow([0]).should eq 0  
end
