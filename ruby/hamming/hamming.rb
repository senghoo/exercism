
class String
  def to_a
    each_char.to_a
  end
end

class Hamming
  VERSION = 1
  def self.compute(a, b)
    fail ArgumentError if a.length != b.length
    a.to_a.zip(b.to_a).count { |x, y| x != y }
  end
end
