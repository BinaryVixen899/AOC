# You come across an experimental new kind of memory stored on an infinite two-dimensional grid.

# Each square on the grid is allocated in a spiral pattern starting at a location marked 1 and then counting up while spiraling outward. 
# For example, the first few squares are allocated like this:
  
# 17  16  15  14  13
# 18   5   4   3  12
# 19   6   1   2  11
# 20   7   8   9  10
# 21  22  23---> ...

# While this is very space-efficient (no squares are skipped), requested data must be carried back to square 1 
# (the location of the only access port for this memory system) by programs that can only move up, down, left, or right. 
# They always take the shortest path: the Manhattan Distance between the location of the data and square 1.

# For example:

# Data from square 1 is carried 0 steps, since it's at the access port.
# Data from square 12 is carried 3 steps, such as: down, left, left.
# Data from square 23 is carried only 2 steps: up twice.
# Data from square 1024 must be carried 31 steps.
# How many steps are required to carry the data from the square identified in your puzzle input all the way to the access port?

# Your puzzle input is 277678.
  
#----------------------------------------------------------------------------------------------------------------------------#

#Class Code not mine
# class Point:
#   def __init__(self, x, y):
#     self.x, self.y = x, y

#   def __str__(self):
#     return "{}, {}".format(self.x, self.y)

#   def __neg__(self):
#     return Point(-self.x, -self.y)

#   def __add__(self, point):
#     return Point(self.x+point.x, self.y+point.y)

#   def __sub__(self, point):
#     return self + -point

# #Variables
# originpoint = Point(3,3)
# currentpoint = Point(3,3) 
# newpointamount = Point(0,0)
# realnumber = 1
# incrementamount = 0
# incrementedamount = 0
# targetnumber = 277678

# #Calculate the new bit
# while realnumber != targetnumber:
#     #SpiralCode 

#     incrementedamount = 0 #reset
#     #I need to find some way to get it to count the inbetween numbers, perhaps more while statements?
#     incrementamount += 1
#     print("test")

#     while incrementamount != incrementedamount: #maybe something like this
#     #move right(Increment one)
#     #so basically keep the increment amount system but also use another system as well. It's also possible we could just abandon the coordinates altogether.
#       currentpoint += Point(1,0)
#       incrementedamount += 1
#       realnumber += 1     
#       print("currentpoint after right move is", currentpoint)

#       if realnumber == targetnumber:
#         print(currentpoint)
#         break

#     if realnumber == targetnumber: #should make these a function, etc.
#         print(currentpoint)
#         break
#     incrementedamount = 0 #reset

#     #move up
#     while incrementamount != incrementedamount:
#       currentpoint = currentpoint + Point(0,1)
#       incrementedamount += 1
#       realnumber += 1
#       print("currentpoint after up move is", currentpoint)

  
#       if realnumber == targetnumber:
#         print("currentpoint and program end",currentpoint)
#         break
        
#     if realnumber == targetnumber: #should make these a function, etc.
#         print(currentpoint)
#         break
#     incrementedamount = 0 #reset

#     incrementedamount = 0

#     incrementamount += 1 #tick more 

#     #move left (increment by 1) #this is the problem. it says while 2 does not = one
#     while incrementamount != incrementedamount:
#       incrementedamount += 1
#       currentpoint = currentpoint + Point(-1,0)
#       realnumber += 1
#       print(currentpoint)

#       if realnumber == targetnumber: #should make these a function, etc.
#         print(currentpoint)
#         break

#     if realnumber == targetnumber: #should make these a function, etc.
#         print(currentpoint)
#         break
#     incrementedamount = 0 #reset

#     #move down 
#     while incrementamount != incrementedamount:
#       incrementedamount += 1
#       currentpoint = currentpoint + Point(0,-1)
#       realnumber += 1

#       if realnumber == targetnumber: #should make these a function, etc.
#         print(currentpoint)
#         break

#     if realnumber == targetnumber: #should make these a function, etc.
#         print(currentpoint)
#         break
    
    

# print("final realnumber is",realnumber)
# print("final Currentpoint is",currentpoint)

#----------------#Thoughts-------------------------#  
#It literally is the manhattan distance, it's (x,y - x,y!)
#So I need to find some way to label these with coordinates
#I wonder if there's some sort of coordinate plane...
#so it goes like over 1, up 1, over two, down two, over three, up three, over four, down four, over five, up five
#so it goes over and up increasing by one, and then over and down increasing by one, and then over and up increasing by one, and then over and down increasing by one, etc.


#------------Advent of Code 3B----------#

# --- Part Two ---
# As a stress test on the system, the programs here clear the grid and then store the value 1 in square 1. 
# Then, in the same allocation order as shown above, they store the sum of the values in all adjacent squares, including diagonals.

# So, the first few squares' values are chosen as follows:

# Square 1 starts with the value 1.
# Square 2 has only one adjacent filled square (with value 1), so it also stores 1.
# Square 3 has both of the above squares as neighbors and stores the sum of their values, 2.
# Square 4 has all three of the aforementioned squares as neighbors and stores the sum of their values, 4.
# Square 5 only has the first and fourth squares as neighbors, so it gets the value 5.
# Once a square is written, its value does not change. Therefore, the first few squares would receive the following values:

# 147  142  133  122   59
# 304    5    4    2   57
# 330   10    1    1   54
# 351   11   23   25   26
# 362  747  806--->   ...
# What is the first value written that is larger than your puzzle input?

#--------------------------------------------------------------------#
#--------------------Thoughts--------------------------------#
#I need to teach it what is adjacent 
#need to determine if we're on the first row or the second row, affects a lot 
#Okay so like, there are 9 squares 
#Maybe I can figure out how to hardcode it in so like, each square adjacent???
#or wait, how many numbers inbetween like, time wise, hmmm  
#hmmm okay so like 1, then 1,
#wait wait wait  square /5/ is the one that breaks the pattern!!!
#okay that's not it but 4 & 10 are in the same relative locations and each use three
#2, 5, and 11 all use 2 each
#except all this falls apart at 54 
#maybe increment by 1?
#no need for grids anymore thank Celestia
#maybe have 1 be say, a center number? 
# and then, so basically "if center #?", right #, etc."
#adjacent #s for (5) list
#FIGURED IT OUT. I NEED TO DEFINE RULES FOR ADJACENCY 
#Okay i'm officially /stuck/. I'm going to go to the next problem. 

# 1
# 1 ? 
# 1 1
#   ?
# 1 1

#  2
#1 1

#  2
#1 1 

# ? 2
# 1 1

# 4 2
# 1 1

# Your puzzle input is still 277678.
#
#

#Variables
currentnumber = 1
centersquare = 1
targetnumber = 277678
incrementamount = 0
incrementedamount = 0

firsttimerun = true
Squareon
ringon


#Calculate the new bit
while currentnumber < targetnumber:
    #SpiralCode 

    incrementedamount = 0 #reset
    incrementamount += 1

    while incrementamount != incrementedamount: #maybe something like this
    #move right(Increment one)
    incrementedamount += 1:

      if firstring(true):
      currentnumber = previous number + centersquare
      

      if secondring(true)
      realnumber += 1     

      if realnumber >= targetnumber:

        break

    if realnumber >= targetnumber: #should make these a function, etc.
        break
    
    incrementedamount = 0 #reset

    #move up
    while incrementamount != incrementedamount:
      incrementedamount += 1

    #if ringone
    if ringone()
    adjacentsquare = 
    previous square 

      if realnumber == targetnumber:
        break
        
    if realnumber == targetnumber: #should make these a function, etc.
        break
    incrementedamount = 0 #reset

    incrementedamount = 0

    incrementamount += 1 #tick more 

    #move left (increment by 1) #this is the problem. it says while 2 does not = one
    while incrementamount != incrementedamount:
      incrementedamount += 1
      realnumber += 1

      if realnumber == targetnumber: #should make these a function, etc.
        break

    if realnumber == targetnumber: #should make these a function, etc.
        break
    incrementedamount = 0 #reset

    #move down 
    while incrementamount != incrementedamount:
      incrementedamount += 1
      realnumber += 1

      if realnumber == targetnumber: #should make these a function, etc.
        break

    if realnumber == targetnumber: #should make these a function, etc.
        break
    
print("final realnumber is",realnumber)

