require "spec"
require "../src/*"

describe("Basic tests") do
  it "With #{["tail", "body", "head"]} should return #{["head", "body", "tail"]}" { (fix_the_meerkat(["tail", "body", "head"]).should eq ["head", "body", "tail"]) }
  it "With #{["tails", "body", "heads"]} should return #{["heads", "body", "tails"]}" { (fix_the_meerkat(["tails", "body", "heads"]).should eq ["heads", "body", "tails"]) }
  it "With #{["bottom", "middle", "top"]} should return #{["top", "middle", "bottom"]}" { (fix_the_meerkat(["bottom", "middle", "top"]).should eq ["top", "middle", "bottom"]) }
  it "With #{["lower legs", "torso", "upper legs"]} should return #{["upper legs", "torso", "lower legs"]}" { (fix_the_meerkat(["lower legs", "torso", "upper legs"]).should eq ["upper legs", "torso", "lower legs"]) }
  it "With #{["ground", "rainbow", "sky"]} should return #{["sky", "rainbow", "ground"]}" { (fix_the_meerkat(["ground", "rainbow", "sky"]).should eq ["sky", "rainbow", "ground"]) }
end
