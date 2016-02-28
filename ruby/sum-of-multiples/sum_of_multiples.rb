class SumOfMultiples
  def self.to(limit)
    new(3, 5).to limit
  end

  def initialize(*args)
    @multiples = args
  end

  def to(limit)
    (2...limit).select { |n| @multiples.any? { |m| n % m == 0 } }.inject(0, :+)
  end
end
