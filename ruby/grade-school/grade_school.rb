class School
  def initialize()
    @students = {}
  end

  def add(name, grade)
    @students[grade] ||= []
    @students[grade] = (@students[grade] << name).sort
  end

  def to_h
    @students.sort.to_h
  end

  def grade(grade)
    @students[grade] || []
  end

  VERSION = 1
end
