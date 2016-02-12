class Fixnum
  VERSION = 1
  ROMAN_DIGITS = {
    1    => 'I',
    4    => 'IV',
    5    => 'V',
    9    => 'IX',
    10   => 'X',
    40   => 'XL',
    50   => 'L',
    90   => 'XC',
    100  => 'C',
    400  => 'CD',
    500  => 'D',
    900  => 'CM',
    1000 => 'M'
  }

  def to_roman
    thiz = self
    res = ''
    ROMAN_DIGITS.keys.reverse_each do |d|
      res += ROMAN_DIGITS[d] * (thiz / d)
      thiz = thiz % d
    end
    res
  end
end
