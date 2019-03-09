require "templates/version"

module Nx
  def self.set(hash, path, value)
    key, sub_key = path.split(".", 2)

    if sub_key.nil?
      hash[key] = value
    else
      if hash[key].nil?
        hash[key] = {}
      end
      self.set(hash[key], sub_key, value)
    end
  end

  def self.get(hash, path)
    result = hash
    path.split(".").each do |dot_part|
      dot_part.split("[").each do |part|
        if part.include?("]")
          index = part.to_i
          result = result[index] rescue nil
        else
          result = result[part] rescue nil
        end
      end
    end
    result
  end

  def self.mix(*args)
    result = {}
    args.each do |arg|
      result.merge! arg
    end
    result
  end
end
