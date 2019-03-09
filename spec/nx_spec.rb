# game_spec.rb
require "nx"

RSpec.describe Nx do
  include Nx
  describe "#set/get with hash" do
    it "set/get with object value" do
      hash = {}
      set(hash, "a.b.c[0]", 12)
      expect(get(hash, "a")).to eq({ "b" => { "c" => [12] } })
    end
  end

  describe "#set/get with []" do
    it "set/get with array value" do
      hash = { "b" => { "c" => [0, [1, 42]] } }
      expect(get(hash, "b.c")).to eq([0, [1, 42]])
      expect(get(hash, "b.c[0]")).to eq(0)
      expect(get(hash, "b.c[1]")).to eq([1, 42])
      expect(get(hash, "b.c[-1]")).to eq([1, 42])
      expect(get(hash, "b.c[-1][0]")).to eq(1)
      expect(get(hash, "b.c[-1][1]")).to eq(42)
    end
  end
end
