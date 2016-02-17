class Fixnum
  def devisible_by?(num)
    self % num == 0
  end
end

class Year
  VERSION = 1
  def self.leap?(y)
    y.devisible_by?(400) || (!y.devisible_by?(100) && y.devisible_by?(4))
  end
end
