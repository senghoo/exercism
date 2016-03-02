class Series
  def initialize(s)
    @series = s.each_char.collect{|i| i.to_i}
  end

  def slices(count)
    fail ArgumentError if count > @series.length
    len = @series.length
    (0..(len - count)).collect { |e| @series[e, count] }
  end
end
