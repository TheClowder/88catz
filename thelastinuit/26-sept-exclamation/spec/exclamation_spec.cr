require "spec"
require "../src/*"

describe "remove_exclamation_marks" do
  remove_exclamation_marks("Hello World!").should eq "Hello World"
  remove_exclamation_marks("Hello World!!!").should eq "Hello World"
  remove_exclamation_marks("Hi! Hello!").should eq "Hi Hello"
  remove_exclamation_marks("").should eq ""
  remove_exclamation_marks("Oh, no!!!").should eq "Oh, no"
end
