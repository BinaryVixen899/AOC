# --- Day 3: Crossed Wires ---
# The gravity assist was successful, and you're well on your way to the Venus refuelling station. During the rush back on Earth, the fuel management system wasn't completely installed, so that's next on the priority list.

# Opening the front panel reveals a jumble of wires. Specifically, two wires are connected to a central port and extend outward on a grid. You trace the path each wire takes as it leaves the central port, one wire per line of text (your puzzle input).

# The wires twist and turn, but the two wires occasionally cross paths. To fix the circuit, you need to find the intersection point closest to the central port. Because the wires are on a grid, use the Manhattan distance for this measurement. While the wires do technically cross right at the central port where they both start, this point does not count, nor does a wire count as crossing with itself.

# Manhattan distance. Definition: The distance between two points measured along axes at right angles. In a plane with p1 at (x1, y1) and p2 at (x2, y2), it is |x1 - x2| + |y1 - y2|. Lm distance.

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

# R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
# U98,R91,D20,R16,D67,R40,U7,R15,U6,R7 = distance 135

# What is the Manhattan distance from the central port to the closest intersection?
#  In a plane with p1 at (x1, y1) and p2 at (x2, y2), it is |x1 - x2| + |y1 - y2|.

#what if we just do the calculations 
#like who cares about the historical points we can add those but
#like right 0,0
#then 75,0 
#then 75 -30
# 
# Jsut need two functions, one that compares these and one that actually does the coordinates  

#{0: [75, 0], 1: [75, -30], 2: [158, -30], 3: [158, 53], 4: [146, 53], 5: [146, 4], 6: [217, 4], 7: [217, 11], 8: [145, 11]}
#{0: [0, 62], 1: [66, 62], 2: [66, 117], 3: [100, 117], 4: [100, 46], 5: [155, 46], 6: [155, -12], 7: [238, -12]}
#WHOOPS THE LOGIC IS WRONg, this does NOT check where they cross because it does NOT look at where they are DURING it just calculates the coordinates
#So what I would have to do would be to like, look at every single coordinate :(
#honestly just use NUMPY


#USE THIS ONE 
# firstwirepath = 'R75,D30,R83,U83,L12,D49,R71,U7,L72'
firstwirepath = 'U62,R66,U55,R34,D71,R55,D58,R83'
pointax = 0 
pointay = 0 
pointbx = 0
pointby = 0
xcord = 0
i = 0
ycord = 0
firstwirepath = firstwirepath.split(',')
i = 0 

dict = {}
print(firstwirepath)

for x in firstwirepath:

    if x[0] == 'R':
        x = x.strip('R')
        xcord += int(x) 
        print(xcord)
        dict[i] = [xcord,ycord]
        print(dict)
        i+=1
        continue
    elif x[0] == 'L': 
        x = x.strip('L')
        xcord += int(x) * -1
        dict[i] = [xcord,ycord]
        i+=1
        continue
    elif x[0] == 'U':
        x = x.strip('U')
        ycord += int(x)
        dict[i] = [xcord,ycord]
        i+=1
        continue
    elif x[0] == 'D':
        x = x.strip('D')
        ycord += int(x) * -1
        dict[i] = [xcord,ycord]
        i+=1
        continue
print(dict)

# for x in firstwirepath:
#     pointax = currentpointa + x
#     pointay = firstwirepath[x+1]
#     currentpointa = [pointax,pointay]

# for x in secondwirepath:
#     pointbx = currentpointb + x
#     pointby = secondwirepath[x+1]
#     currentpointb = [pointbx,pointby]


# if (currentpointa == currentpointb):
#     xdistance = pointax - pointay 
#     ydistane
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
