module Nx
  module HashExt
    def set(path, value)
      Nx.set(self, path, value)
    end

    def get(path)
      Nx.get(self, path)
    end

    def mix(*args)
      Nx.mix(self, *args)
    end
  end
end
