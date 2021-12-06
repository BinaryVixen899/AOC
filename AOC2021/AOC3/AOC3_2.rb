

# Both the oxygen generator rating and the CO2 scrubber rating are values that can be found in your diagnostic report - finding them is the tricky part. Both values are located using a similar process that involves filtering out values until only one remains. Before searching for either rating value, start with the full list of binary numbers from your diagnostic report and consider just the first bit of those numbers. Then:

# Keep only numbers selected by the bit criteria for the type of rating value for which you are searching. Discard numbers which do not match the bit criteria.
# If you only have one number left, stop; this is the rating value for which you are searching.
# Otherwise, repeat the process, considering the next bit to the right.
# The bit criteria depends on which type of rating value you want to find:

# To find oxygen generator rating, determine the most common value (0 or 1) in the current bit position, and keep only numbers with that bit in that position. If 0 and 1 are equally common, keep values with a 1 in the position being considered.

# To find CO2 scrubber rating, determine the least common value (0 or 1) in the current bit position, and keep only numbers with that bit in that position. If 0 and 1 are equally common, keep values with a 0 in the position being considered.

# For example, to determine the oxygen generator rating value using the same example diagnostic report from above:

# Start with all 12 numbers and consider only the first bit of each number. There are more 1 bits (7) than 0 bits (5), so keep only the 7 numbers with a 1 in the first position: 11110, 10110, 10111, 10101, 11100, 10000, and 11001.


# Then, consider the second bit of the 7 remaining numbers: there are more 0 bits (4) than 1 bits (3), so keep only the 4 numbers with a 0 in the second position: 10110, 10111, 10101, and 10000.
# In the third position, three of the four numbers have a 1, so keep those three: 10110, 10111, and 10101.
# In the fourth position, two of the three numbers have a 1, so keep those two: 10110 and 10111.
# In the fifth position, there are an equal number of 0 bits and 1 bits (one each). So, to find the oxygen generator rating, keep the number with a 1 in that position: 10111.
# As there is only one number left, stop; the oxygen generator rating is 10111, or 23 in decimal.
# Then, to determine the CO2 scrubber rating value from the same example above:

# Start again with all 12 numbers and consider only the first bit of each number. There are fewer 0 bits (5) than 1 bits (7), so keep only the 5 numbers with a 0 in the first position: 00100, 01111, 00111, 00010, and 01010.
# Then, consider the second bit of the 5 remaining numbers: there are fewer 1 bits (2) than 0 bits (3), so keep only the 2 numbers with a 1 in the second position: 01111 and 01010.
# In the third position, there are an equal number of 0 bits and 1 bits (one each). So, to find the CO2 scrubber rating, keep the number with a 0 in that position: 01010.
# As there is only one number left, stop; the CO2 scrubber rating is 01010, or 10 in decimal.
# Finally, to find the life support rating, multiply the oxygen generator rating (23) by the CO2 scrubber rating (10) to get 230.

# Use the binary numbers in your diagnostic report to calculate the oxygen generator rating and CO2 scrubber rating, then multiply them together. What is the life support rating of the submarine? (Be sure to represent your answer in decimal, not binary.)
require 'pry'


$data = [
  '00100',
  '11110',
  '10110',
  '10111',
  '10101',
  '01111',
  '00111',
  '11100',
  '10000',
  '11001',
  '00010',
  '01010']

# def oxygeneratorrating

# lifesupportrating = oxygeneratorrating * CO2scrubberrating 

# oxygeneratorrating = oxygeneratorratingcalculation

# $data = File.readlines("$data.txt")
# $data = $data.map {|x| x.to_s}

$zerocount = 0 
$onecount = 0 
$round = 0 
length = $data[0].length
$originaldata = $data.clone
binding.pry
def increment_one
  $onecount+= 1 
end

def increment_zero
  $zerocount += 1
end

def calculate_oxygeon
  if $onecount > $zerocount
    $data.reject! { |x| x[$round].to_i != 1 }
    # $data.map!.with_index{ |x, i| x[$round].to_i != 1 ? $data.delete(x) : ""} I want to understand why this fails 
  elsif $zerocount > $onecount
    $data.reject! { |x| x[$round].to_i != 0 }
  else
    $data.reject! { |x| x[$round].to_i != 1 }
  end

  # if $onecount > $zerocount
  #   # for i in $data 
  #   #   if $round == 0
  #   #     puts i 
  #   #   end
  #     $data.map!(i[$round].to_i != 1 ? $data.delete(i) : "")
  #     # we're skipping numbers
  #     # Because we're restarting at "for i in we restart at the next one "
  #     #possible we're hitting the 1st one
  #     #so we need to do this with map perhaps
  #     puts "#{$data}"       
  #   end
    
  # elsif 
    
  #   

  # else    
  #   $data.map!(i[$round].to_i != 1 ? $data.delete(i) : "")
  # end
# I could turn all this into a case statement 
end

def calculate_co2
  if $onecount > $zerocount
    $data.reject! { |x| x[$round].to_i != 0 }

    
  elsif $zerocount > $onecount
      $data.reject! { |x| x[$round].to_i != 1 }

  else    
    $data.reject! { |x| x[$round].to_i != 0 }
  end
# I could turn all this into a case statement 
end




while $data.count > 1
  for i in $data
    # puts i[$round]
    i[$round].to_i == 1 ? increment_one() : increment_zero()
    # Turn this into map 
  end
  puts "onecount is #{$onecount}"
  puts "zerocount is #{$zerocount}"
  # In first onecount is 7 and zerocount is 5, so we should be fine here
  puts "#{$round} and #{$data}"
  calculate_oxygeon()
  puts "survivors after #{$round} + #{$data}"
  $data.reject!(&:empty?)
  # 0 and ["11110", "10110", "10111", "10101", "00111", "11100", "10000", "11001", "01010"]
  # First round result from my data, but we should see everything with a 0 eliminated... What went wrong? 

  $onecount = 0
  $zerocount = 0
  $round+= 1 
end

oxygeneratorrating = $data 
puts oxygeneratorrating
puts "this should be 10111"
$data = $originaldata
$round = 0
# As there is only one number left, stop; the oxygen generator rating is 10111, or 23 in decimal.

# co2 calculation
binding.pry

while $data.count > 1
  for i in $data
    i[$round].to_i == 1 ? increment_one() : increment_zero()
  end
  calculate_co2
  puts "#{$round} and #{$data}"
  $onecount = 0
  $zerocount = 0
  $round+= 1 
end

co2rating = $data
puts co2rating 
p epsilonarray.to_i(2) * gammaarray.to_i(2)
puts "this should be 01010 "
# Do bit blasting 
