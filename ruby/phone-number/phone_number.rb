class PhoneNumber
  attr_reader :number
  def initialize(phone)
    @number = '0000000000'
    return if phone.each_char.any? { |e| ('a'..'z').cover? e }
    number = phone.each_char.collect { |e| e if ('0'..'9').cover? e }.join
    if number.length == 10
      @number = number
    elsif number.length == 11 && number[0] == '1'
      @number = number[1..10]
    end
  end

  def to_s
    "(#{@number[0..2]}) #{@number[3..5]}-#{@number[6..10]}"
  end

  def area_code
    @number[0..2]
  end
end
