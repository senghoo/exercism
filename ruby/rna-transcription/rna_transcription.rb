class Complement
  VERSION = 3
  RNA_MAP = { 'G' => 'C',
              'C' => 'G',
              'T' => 'A',
              'A' => 'U' }
  def self.of_dna(s)
    fail ArgumentError if s.each_char.any? { |e| RNA_MAP[e].nil? }
    s.each_char.collect { |e| RNA_MAP[e] }.join
  end
end
