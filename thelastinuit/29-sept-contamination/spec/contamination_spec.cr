require "spec"
require "../src/*"

describe "Basic tests" do
  contamination("abc","z").should eq "zzz"
  contamination("","z").should eq ""
  contamination("abc","").should eq ""
  contamination("_3ebzgh4","&").should eq "&&&&&&&&"
  contamination("//case"," ").should eq "      "
end
