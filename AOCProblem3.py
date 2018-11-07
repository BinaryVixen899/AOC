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

#Generate the Equation
#1+1

#Class Code not mine
class Point:
  def __init__(self, x, y):
    self.x, self.y = x, y

  def __str__(self):
    return "{}, {}".format(self.x, self.y)

  def __neg__(self):
    return Point(-self.x, -self.y)

  def __add__(self, point):
    return Point(self.x+point.x, self.y+point.y)

  def __sub__(self, point):
    return self + -point

#Variables
originpoint = Point(3,3)
currentpoint = Point(3,3) 
newpointamount = Point(0,0)
realnumber = 1
incrementamount = 0
incrementedamount = 0
targetnumber = 2

#Calculate the new bit
while realnumber != targetnumber:
    #SpiralCode

    #I need to find some way to get it to count the inbetween numbers, perhaps more while statements?
    incrementamount += 1

    while incrementamount <= incrementedamount: #maybe something like this
    #move right(Increment one)
    #so basically keep the increment amount system but also use another system as well. It's also possible we could just abandon the coordinates altogether.
      currentpoint += Point(1,0)
      incrementedamount += 1
      realnumber += 1     
      print(currentpoint)

    if realnumber == targetnumber:
      print(currentpoint)
      break
  
    incrementedamount = 0 #reset
    
    #move up
    while incrementamount <= incrementedamount:
      currentpoint = currentpoint + Point(0,1)
      incrementedamount += 1
      realnumber += 1
      print(currentpoint)


    if realnumber == targetnumber:
      print(currentpoint)
      break

    incrementedamount = 0

    incrementamount += 1 #tick more 

    #move left (increment by 1)
    while incrementamount <= incrementedamount:
      incrementedamount += 1
      currentpoint = currentpoint + Point(-1,0)
      realnumber += 1
      print(currentpoint)

    if realnumber == targetnumber: #should make these a function, etc.
      print(currentpoint)
      break

    incrementedamount = 0 #reset

    #move down 
    while incrementamount <= incrementedamount:
      incrementedamount += 1
      currentpoint = currentpoint + Point(0,-1)
      realnumber += 1

    if realnumber == targetnumber: #should make these a function, etc.
      print(currentpoint)
      break
    
print("realnumber",realnumber)
print("Currentpoint",currentpoint)

#----------------#Thoughts-------------------------
#It literally is the manhattan distance, it's (x,y - x,y!)
#So I need to find some way to label these with coordinates
#I wonder if there's some sort of coordinate plane...
#so it goes like over 1, up 1, over two, down two, over three, up three, over four, down four, over five, up five
#so it goes over and up increasing by one, and then over and down increasing by one, and then over and up increasing by one, and then over and down increasing by one, etc.
# 