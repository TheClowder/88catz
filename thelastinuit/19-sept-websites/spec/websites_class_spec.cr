require "spec"
require "../src/*"


describe "Tests" do
  it "Testing if the item has size 1000" do websites.size.should eq 1000 end
  it "Testing if the item has class 'Array'" do websites.class.should eq Array(String) end
  it "Testing if the items contains only 'codewars'" do websites.to_set.should eq Set{"codewars"} end
end
