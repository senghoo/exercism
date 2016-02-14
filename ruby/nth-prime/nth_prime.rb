class Prime
  def self.nth(n)
    fail ArgumentError if n <= 0
    lst = [2]
    i = 2
    until lst.length >= n
      i += 1
      next if lst.count { |j| i % j == 0 } > 0
      lst << i
    end
    lst[-1]
  end
end
