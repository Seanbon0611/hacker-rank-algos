def compareTriplets(a, b)
    alices_score = 0
    bobs_score = 0
   a.each_with_index do |v, i|
    if v > b[i]
        alices_score += 1
    elsif v < b[i]
        bobs_score += 1
    end
   end
   return alices_score, bobs_score
end