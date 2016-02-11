class Squares
  VERSION = 2
  def initialize(n)
    @n = n
  end

  def square_of_sum
    ((1 + @n) * @n / 2)**2
  end

  def sum_of_squares
    @n * (@n + 1) * (2 * @n + 1) / 6
  end

  def difference
    square_of_sum - sum_of_squares
  end
end
