class Raindrops
  VERSION = 1
  MAP = { 3 => 'Pling',
          5 => 'Plang',
          7 => 'Plong' }
  def self.convert(v)
    r = [3, 5, 7].each.collect { |e|
      MAP[e] if v % e == 0
    }.join
    r == '' ? v.to_s : r
  end
end
