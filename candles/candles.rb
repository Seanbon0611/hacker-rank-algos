def birthdayCakeCandles(candles)
  map = {}
  candles.each_with_index do |num, i|
      if map[candles[i]]
          map[candles[i]] += 1
      else
          map[candles[i]] = 1
      end
  end
  map.each do |k, v| 
      return map.values.max
  end
end

birthdayCakeCandles([3,1,2,3])