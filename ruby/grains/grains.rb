class Grains
  def self.square(n)
    1 << (n-1)
  end
  def self.total
    (1..64).collect { |e| square e }.inject(0, :+)
  end
end
