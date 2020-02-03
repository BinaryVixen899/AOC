# --- Day 3: Crossed Wires ---
# The gravity assist was successful, and you're well on your way to the Venus refuelling station. During the rush back on Earth, the fuel management system wasn't completely installed, so that's next on the priority list.

# Opening the front panel reveals a jumble of wires. Specifically, two wires are connected to a central port and extend outward on a grid. You trace the path each wire takes as it leaves the central port, one wire per line of text (your puzzle input).

# The wires twist and turn, but the two wires occasionally cross paths. To fix the circuit, you need to find the intersection point closest to the central port. Because the wires are on a grid, use the Manhattan distance for this measurement. While the wires do technically cross right at the central port where they both start, this point does not count, nor does a wire count as crossing with itself.

# For example, if the first wire's path is R8,U5,L5,D3, then starting from the central port (o), it goes right 8, up 5, left 5, and finally down 3:

# ...........
# ...........
# ...........
# ....+----+.
# ....|....|.
# ....|....|.
# ....|....|.
# .........|.
# .o-------+.
# ...........
# Then, if the second wire's path is U7,R6,D4,L4, it goes up 7, right 6, down 4, and left 4:

# ...........
# .+-----+...
# .|.....|...
# .|..+--X-+.
# .|..|..|.|.
# .|.-X--+.|.
# .|..|....|.
# .|.......|.
# .o-------+.
# ...........
# These wires cross at two locations (marked X), but the lower-left one is closer to the central port: its distance is 3 + 3 = 6.

# Here are a few more examples:

# R75,D30,R83,U83,L12,D49,R71,U7,L72
# U62,R66,U55,R34,D71,R55,D58,R83 = distance 159
<<<<<<< HEAD

=======
>>>>>>> 08074ab64e43f649258c661a56a1c24687b4b307
# R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
# U98,R91,D20,R16,D67,R40,U7,R15,U6,R7 = distance 135
# What is the Manhattan distance from the central port to the closest intersection?
#  In a plane with p1 at (x1, y1) and p2 at (x2, y2), it is |x1 - x2| + |y1 - y2|.

<<<<<<< HEAD
firstwirepath = [R75,D30,R83,U83,L12,D49,R71,U7,L72]
secondwirepath = [U62,R66,U55,R34,D71,R55,D58,R83]
firstwirepathrefined=[]
secondwirepathrefined=[]
pointax = 0 
pointay = 0 
pointbx = 0
pointby = 0
currentpointa = 0
currentpointb = 0

for x in firstwirepath:
    if x[1] == R:
        x[1]
        firstwirepathrefined.append(x)
    elif x[1] == L: 
        firstwirepathrefined.append(-x[1])
    elif x[1] == U:
        firstwirepathrefined.append(x)
    elif x[1] == D:
        firstwirepathrefined.append(-x[1])

for x in firstwirepath:
    pointax = currentpointa + x
    pointay = firstwirepath[x+1]
    currentpointa = [pointax,pointay]

for x in secondwirepath:
    pointbx = currentpointb + x
    pointby = secondwirepath[x+1]
    currentpointb = [pointbx,pointby]


if (currentpointa == currentpointb):
    xdistance = pointax - pointay 
    ydistane
=======
# firstwirepath = ['R75','D30','R83','U83','L12','D49','R71','U7','L72']
# secondwirepath = ['U62','R66','U55','R34','D71','R55','D58','R83']

# firstwirepathworkingcoordinates = []
# currentxcoordinate = 0
# currentycoordinate = 0

# for x in firstwirepath:
#     if x[0] == 'L':
#         x = x.strip('L')
#         firstwirepathworkingcoordinates.append(int(x) * -1)
#         currentxcoordinate = currentxcoordinate + x
#         exit()
#     elif x[0] == 'R':
#         x = x.strip('R')
#         currentxcoordinate = currentxcoordinate + x lambdx
#         firstwirepathworkingcoordinates.append(x)
#     elif x[0] == 'D':
#         x = x.strip('D')
#         firstwirepathworkingcoordinates.append(x)

    
#     if x[0] == 'U':
#         pass
#     if x[0] == 'D':
# pointa1 = 0
# pointa2 = 0 
# pointb1 = 0 
# pointb2 = 0

# for x in firstwirepath:
#     if

# if (pointa1  == point b1 and pointa2 ==b2):
#     crossoverpoints.append[pointa1]
>>>>>>> 08074ab64e43f649258c661a56a1c24687b4b307
