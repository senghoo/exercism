class Prime
  def self.nth(n)
    fail ArgumentError if n <= 0
    lst = [2]
    i = 2
    until lst.length == n
      i += 1
      sqrt_i = Math.sqrt i
      next if lst.find { |prime|
        break nil if prime > sqrt_i
        i % prime == 0
      }
      lst << i
    end
    lst.last
  end
end
