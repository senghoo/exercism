class Prime
  def self.nth(n)
    fail ArgumentError if n <= 0
    lst = [2]
    i = 2
    until lst.length == n
      i += 1
      lst << i if lst.all? { |prime| i % prime != 0 }
    end
    lst.last
  end
end
