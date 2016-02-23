class Sieve
  def initialize(max)
    @max = max
  end

  def primes
    res = (2..@max).to_a
    sqrt_of_max = Math.sqrt @max
    res.each do |prime|
      return res if prime > sqrt_of_max
      res.delete_if do |num|
        d, m = num.divmod prime
        d != 1 && m == 0
      end
    end
  end
end
