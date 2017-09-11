require "spec"
require "../src/*"

describe "Basic tests" do
  it "With #{[9, 3, '7', '3']} should return 22" { sum_mix([9, 3, '7', '3']).should eq 22} 
  it "With #{['5', '0', 9, 3, 2, 1, '9', 6, 7]} should return 42" { sum_mix(['5', '0', 9, 3, 2, 1, '9', 6, 7]).should eq 42} 
  it "With #{['3', 6, 6, 0, '5', 8, 5, '6', 2,'0']} should return 41" { sum_mix(['3', 6, 6, 0, '5', 8, 5, '6', 2,'0']).should eq 41} 
  it "With #{['1', '5', '8', 8, 9, 9, 2, '3']} should return 45" { sum_mix(['1', '5', '8', 8, 9, 9, 2, '3']).should eq 45} 
  it "With #{[8, 0, 0, 8, 5, 7, 2, 3, 7, 8, 6, 7]} should return 61" { sum_mix([8, 0, 0, 8, 5, 7, 2, 3, 7, 8, 6, 7]).should eq 61} 
end
