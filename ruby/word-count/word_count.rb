class Phrase
  VERSION = 1
  def initialize(str)
    @str = str
  end

  def word_count
    @str.split(/'?[^a-zA-Z0-9_']+'?/).group_by{|e| e.downcase}.map { |w, ws| [w, ws.length] }.to_h
  end
end
