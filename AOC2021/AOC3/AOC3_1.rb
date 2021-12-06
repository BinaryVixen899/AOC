# powerconsumption = gamma_rate * epislon_rate
# gr = most common bit in the corresponding position

# Considering only the first bit of each number, there are five 0 bits and seven 1 bits. Since the most common bit is 1, the first bit of the gamma rate is 1.

# The most common second bit of the numbers in the diagnostic report is 0, so the second bit of the gamma rate is 0.

# The most common value of the third, fourth, and fifth bits are 1, 1, and 0, respectively, and so the final three bits of the gamma rate are 110.

# So, the gamma rate is the binary number 10110, or 22 in decimal.

# The epsilon rate is calculated in a similar way; rather than use the most common bit, the least common bit from each position is used. So, the epsilon rate is 01001, or 9 in decimal. Multiplying the gamma rate (22) by the epsilon rate (9) produces the power consumption, 198.
gammaarray = []
epsilonarray = []

data = [
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
data = File.readlines("data.txt")
data = data.map {|x| x.to_s}
$zerocount = 0 
$onecount = 0 
round = 0 

def increment_one
  $onecount+= 1 
end

def increment_zero
  $zerocount += 1
end

def calculateendofround (gammaarray, epsilonarray)
  $onecount > $zerocount ? gammaarray.push(1) && epsilonarray.push(0) : gammaarray.push(0) && epsilonarray.push(1)
end

while round < 12
  for i in data
    # puts i[round]
    i[round].to_i == 1 ? increment_one() : increment_zero()
  end
  calculateendofround(gammaarray, epsilonarray)
  $onecount = 0
  $zerocount = 0
  round+= 1 
end

gammaarray = gammaarray.each {|x| x.to_s}
p gammaarray = gammaarray.join

# puts epsilonarray

epsilonarray = epsilonarray.each {|x| x.to_s}
epsilonarray = epsilonarray.join
p epsilonarray.to_i(2) * gammaarray.to_i(2)
#lets do 
