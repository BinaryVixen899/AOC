 numbers = [7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1]
#numbers = [69,88,67,56,53,97,46,29,37,51,3,93,92,78,41,22,45,66,13,82,2,7,52,40,18,70,32,95,89,64,84,68,83,26,43,0,61,36,57,34,80,39,6,63,72,98,21,54,23,28,65,16,76,11,20,33,96,4,10,25,30,19,90,24,55,91,15,8,71,99,58,14,60,48,44,17,47,85,74,87,86,27,42,38,81,79,94,73,12,5,77,35,9,62,50,31,49,59,75,1]
require './bingocard.rb'
require 'pry'
#Okay so we have rows and columns 
#I probably want to use pointers 
#I need a thing for detecting if numbers are next to each other
#Liz was right I'm going to have to use customer classes
#This is going to be fun!



myarray = []
mycards = []
a = IO.readlines("testbingocards.txt", chomp: true)
a.each do |line|
   myarray.push(line.split())
end
puts myarray
myarray.each {|x| x.to_s} && myarray.flatten! && myarray.map!(&:to_i)
puts myarray
myarray.each_slice(25).with_index {|(*a), i| mycards = mycards.push(BingoCard.new("BingoCard#{i}", a, numbers))}
mycards.each {|x| x.execute}

mycards.each{|a| puts "lastnumber is #{a.lastnumber} and score is #{a.sum} and ttw is #{a.turnstowin} and called numbers is #{a.calledarray} and solveable array is #{a.solvabledictionary.values}"  }

  # a.each  {|a| myarray.push(a) }
# puts "My array is #{myarray}"
# myarray.each do |sub_array|
#   sub_array.each do |item|
#     item.each_char do |number| 
#       p number
#     end

    # item.each do |number|
    #   p number
    # end

    # item.each do |number| 
    #   p number
#     # end
#   end
# end
# after every number drawn, check the adjacent numbers 
# then check if any of those adjacent numbers are a match to the numbers we have already drawn 
# if they are a match then add that to some sort of match tracker 
# etc. 
#69,88,67,56,53,97,46,29,37,51,3,93,92,78,41,22,45,66,13,82,2,7,52,40,18,70,32,95,89,64,84,68,83,26,43,0,61,36,57,34,80,39,6,63,72,98,21,54,23,28,65,16,76,11,20,33,96,4,10,25,30,19,90,24,55,91,15,8,71,99,58,14,60,48,44,17,47,85,74,87,86,27,42,38,81,79,94,73,12,5,77,35,9,62,50,31,49,59,75,1

# 