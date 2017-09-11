require "spec"
require "../src/*"

describe "hydrate!" do
  litres(2).should eq 1
  litres(1).should eq 0
  litres(10).should eq 5
  litres(1.4).should eq 0
  litres(12.3).should eq 6
  litres(0.82).should eq 0
  litres(11.8).should eq 5
  litres(1787).should eq 893
  litres(0).should eq 0
  litres(5.5).should eq 2
end
