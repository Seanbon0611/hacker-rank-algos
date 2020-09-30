my_array = [[1,3,9], [7,4,2], [5,2,4]]

def diagonalDifference(arr)
  left_diag_sum = 0
  right_diag_sum = 0
  arr.each_with_index do |ar, i|
    #iterates through each array within nested array and gets the value assigned to that index
    left_diag_sum += ar[i]
    #does the same but when you add the '-' in front of the i it will go backwards so we can go diagonal from the right
    right_diag_sum += ar[-i-1]
  end
  #abs returns the absolute value of a number
(right_diag_sum - left_diag_sum).abs
end