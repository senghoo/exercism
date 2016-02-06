
class String
  def to_a
    each_char.to_a
  end
end

class Hamming
  def self.compute(a, b)
    a.to_a.zip(b.to_a).collect { |j, k| j == k ? 0 : 1}.inject(:+)
  end
end
