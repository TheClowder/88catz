require "spec"
require "../src/*"

describe "Basic reverses" do
  reverse("I am an expert at this").should eq "this at expert an am I"
  reverse("This is so easy").should eq "easy so is This"
  reverse("no one cares").should eq "cares one no"
  reverse("").should eq ""
  reverse("CodeWars").should eq "CodeWars"
end
