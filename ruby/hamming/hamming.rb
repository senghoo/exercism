
class String
  def to_a
    each_char.to_a
  end
end

class Hamming
  VERSION = 1
  def self.compute(a, b)
    fail ArgumentError if a.length != b.length
    a.to_a.zip(b.to_a).collect { |j, k| j == k ? 0 : 1}.inject(:+) || 0
  end
end
