def mini_min_max(arr)
  sorted = arr.sort! {|a,b| a <=> b}
  min = 0
  max = 0
  for i in 0...(sorted.length - 1) do
    min += sorted[i]
    max += sorted[i + 1]
  end
  print "#{min} #{max}"
end

mini_min_max([1,2,3,4,5])