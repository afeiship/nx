require "templates/version"

module Nx
  def set(path, value)
    me = self || {}
    keys = []
    path.split(".").each do |dot_part|
      dot_part.split("[").each do |part|
        keys << (part.include?("]") ? part.to_i : part)
      end
    end

    # set by path:
    keys[0..-2].each_with_index do |key, index|
      the_key = key
      the_key = yeild(the_key) if block_given?
      me = me[the_key] = me[the_key] || (keys[index + 1].class == String ? {} : [])
    end

    the_key = keys[-1]
    the_key = yeild(the_key) if block_given?
    me[the_key] = value
  end

  def get(path)
    result = self
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

# extend ruby nx
class Hash
  include Nx
end
