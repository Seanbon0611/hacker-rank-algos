a = [4,22,6]
b = [1,22,20]

def compareTriplets(a, b)
    alices_score = 0
    bobs_score = 0
   a.each_with_index do |v, i|
    #if the value is greater than the value stored at b array's index
    if v > b[i]
        alices_score += 1
    elsif v < b[i]
        bobs_score += 1
    end
   end
   return alices_score, bobs_score
end