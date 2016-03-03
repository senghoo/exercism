class Series
  def initialize(s)
    @series = s.each_char.collect &:to_i
  end

  def slices(size)
    Array.new(slices_count size) { |e| @series[e, size] }
  end

  def slices_count(size)
    len = @series.length
    fail ArgumentError if size > len
    len - size + 1
  end
end
