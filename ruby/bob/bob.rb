class Bob
  def hey(remark)

    if remark.upcase == remark && remark.downcase != remark
      'Whoa, chill out!'
    elsif remark.end_with?('?')
      "Sure."
    elsif remark !~ /[^[:space:]]/
      'Fine. Be that way!'
    elsif remark.end_with? '?'
      "Sure."
    else
      'Whatever.'
    end
  end
end
