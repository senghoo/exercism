class Binary
  VERSION = 1
  def initialize(bin)
    @bin = bin
    fail ArgumentError if bin.each_char.any? {|c| !["0", "1"].include?(c)}
  end

  def to_decimal
    res = 0
    @bin.each_char do |char|
      res <<= 1
      res += 1 if char == '1'
    end
    res
  end
end
