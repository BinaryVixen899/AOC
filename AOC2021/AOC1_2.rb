# okay so each new number increments a new letter and consists of the next two numbers
# I could do this with an array of arrays 
# I really only need two sets of arrays 
# stor the initial array have your thing be able to operate on arrays or strings of data and have something that transforms your 2000 in 1997 instead
# taking the chunks and creating something that's a modified array instead 
#instea dof recomputing the rolling sums as you go along, keep track of the most recent three individuals
#1
#2 
#3 
#4 
#5

# 6 (sum of first 3)
# 6 (subtract 1 add 4)
# require 'IO'
#subtract the 1st item in the last array 
#add the first item of the new array 
# text = File.open('data.txt').read
require 'pry'
increasecount = 0
previousvalue = 0
i = 0 

array = [
199,
200,
208,
210,
200,
207,
240,
269,
260,
263
]

array = File.readlines("data.txt", sep=$/)
array.map!(&:to_i)
binding.pry
def getsum(anarray)
  for i in anarray
    puts i
  end
end



while i < array.length
  arraya = array[i, 3]
  puts arraya
  # array a calculation and sum goes here
  i += 1 
  # binding.pry
  arrayb = arraya[1,2].push(array[i+2])
  puts arrayb
  puts increasecount
  # binding.pry
  if arrayb.sum > arraya.sum
    increasecount += 1
  end
  puts "increasecount is #{increasecount}"
  puts "arraya is #{arraya}"
  puts "arrayb is #{arrayb}"
  # check if the last three items in array b are equal to the last three items of the list, if they are then let's end this
  puts "i is #{i}"
  if i + 3 > array.length
    puts "end of the line"
    exit
  end

end



  # array b calculation and sum goes here 

# File.foreach('data.txt').with_index do |line, line_no|
#   currentvalue = Integer(line)
#   if line_no != 0 
#     if currentvalue > previousvalue
#       increasecount += 1    
#       previousvalue = currentvalue
#     else
#       previousvalue = currentvalue
#     end

#   else
#      zerocount +=1
#     previousvalue = currentvalue  
#   end
  
#   puts "increase count: #{increasecount}"
# end

# while arraya.len < 2
#   arraya.push(line)

# while arrayb < 2
#   arrayb.push(line)