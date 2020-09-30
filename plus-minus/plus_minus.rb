array = [1,1,0,-1,-1]

def plusMinus(arr)
  positives = 0
  negatives = 0
  zeros = 0
  arr.each do |num|
    if num < 0
      negatives += 1
    elsif num == 0
      zeros += 1
    else
      positives += 1
    end
  end
  #sprintf Returns the string resulting from applying format_string to any additional arguments. 
  #this allows us to ensure that there are 6 characters after the decimal
    puts sprintf("%6f", positives.fdiv(arr.length))
    puts sprintf("%6f", negatives.fdiv(arr.length))
    puts sprintf("%6f", zeros.fdiv(arr.length))
end