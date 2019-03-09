require "templates/version"

module Nx
  def set(hash, path, value)
    me = hash || {}
    keys = []
    path.split(".").each do |dot_part|
      dot_part.split("[").each do |part|
        keys << (part.include?("]") ? part.to_i : part)
      end
    end
    # set by path:
    keys[0..-2].each do |key|
      me = me[key] = me[key] || (key.class == String ? {} : [])
    end
    me[keys[-1]] = value
  end

  def get(hash, path)
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

  def mix(*args)
    result = {}
    args.each do |arg|
      result.merge! arg
    end
    result
  end
end
