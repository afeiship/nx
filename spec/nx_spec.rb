# game_spec.rb
require "gem-next-core"

RSpec.describe Nx do
  describe "#set/get with hash" do
    it "set/get with object value" do
      #   game = Game.new
      #   20.times { game.roll(0) }
      #   expect(game.score).to eq(0)
      hash = {}
      Nx.set(hash, "a.b", 1)
      expect(Nx.get(hash, "a")).to eq({ "b" => 1 })
    end
  end

  describe "#set/get with []" do
    it "set/get with array value" do
      hash = { "b" => { "c" => [0, [1, 42]] } }
      expect(Nx.get(hash, "b.c")).to eq([0, [1, 42]])
      expect(Nx.get(hash, "b.c[0]")).to eq(0)
      expect(Nx.get(hash, "b.c[1]")).to eq([1, 42])
      expect(Nx.get(hash, "b.c[-1]")).to eq([1, 42])
      expect(Nx.get(hash, "b.c[-1][0]")).to eq(1)
      expect(Nx.get(hash, "b.c[-1][1]")).to eq(42)
    end
  end
end
