class Robot
  attr_reader :name
  def initialize
    @history = []
    reset
  end

  def reset
    @history << @name
    @name = new_name while @history.include? @name
  end

  private

  def new_name
    get_rand(2, ('A'..'Z').to_a) + get_rand(3, ('0'..'9').to_a)
  end

  def get_rand(count, array)
    Array.new(count){ array.to_a[rand(array.length)] }.join
  end
end
