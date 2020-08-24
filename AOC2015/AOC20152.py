# --- Day 2: I Was Told There Would Be No Math ---
# The elves are running low on wrapping paper, and so they need to submit an order for more. They have a list of the dimensions (length l, width w, and height h) of each present, and only want to order exactly as much as they need.

# Fortunately, every present is a box (a perfect right rectangular prism), which makes calculating the required wrapping paper for each gift a little easier: find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l. The elves also need a little extra paper for each present: the area of the smallest side.

# For example:

# A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
# A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.
# All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?
import os

# list = []
# value = 0
# try:
#     os.chdir(r"C:\Users\Alexandria Leal\Documents\Coding\AOC\AOC2015")
# except:
#     #If you're wondering what hapepned here, I decided at some point to just code on my Ubuntu laptop for, who knows waht reason, so I went and di that
#     os.chdir('./AOC/AOC2015/')
#     #In the real world I would, and you should, use an absolute path or such but well you know, I don't want my username being shown in this for Reasons


# with open('AOC20152Input','r') as f:
#     for lin in f:
#         list.append(lin.strip('\n'))

# print(list)

# def MainFunction(l,w,h):
#     a = l*w
#     b = w*h
#     c = h*l
#     return (2 * a) + (2 * b) + (2 * c) + min(a,b,c)



# input = ['2x3x4']

# for i in list:
#     testinput = i.split('x')
#     value += (MainFunction(int(testinput[0]),int(testinput[1]),int(testinput[2])))

# print(value)

# for i in testinput:
#     l = i[a]

# 2*l*w + 2*w*h + 2*h*l

#---------------------------------------------
#Problem 2
# --- Part Two ---
# The elves are also running low on ribbon. Ribbon is all the same width, so they only have to worry about the length they need to order, which they would again like to be exact.

# The ribbon required to wrap a present is the shortest distance around its sides, or the smallest perimeter of any one face. Each present also requires a bow made out of ribbon as well; the feet of ribbon required for the perfect bow is equal to the cubic feet of volume of the present. Don't ask how they tie the bow, though; they'll never tell.

# For example:

# A present with dimensions 2x3x4 requires 2+2+3+3 = 10 feet of ribbon to wrap the present plus 2*3*4 = 24 feet of ribbon for the bow, for a total of 34 feet.
# A present with dimensions 1x1x10 requires 1+1+1+1 = 4 feet of ribbon to wrap the present plus 1*1*10 = 10 feet of ribbon for the bow, for a total of 14 feet.
# How many total feet of ribbon should they order?

# Although it hasn't changed, you can still get your puzzle input.


mylist = []
value = 0
y = 0
x = 0


def Prob2 (
    parameter_list):
    pass
try:
    os.chdir(r"C:\Users\Alexandria Leal\Documents\Coding\AOC\AOC2015")
except:
    #If you're wondering what hapepned here, I decided at some point to just code on my Ubuntu laptop for, who knows waht reason, so I went and di that
    os.chdir('./AOC/AOC2015/')
    #In the real world I would, and you should, use an absolute path or such but well you know, I don't want my username being shown in this for Reasons

with open('AOC20152Input','r') as f:
    for lin in f:
        mylist.append(lin.strip('\n'))

# mylist = list(map(list, mylist))
print(str(mylist))
newlist =[]
[newlist.append(a) for a in [i.split('x') for i in mylist]]
print(newlist)
for a in newlist:
    print(a[1])
    



#     y += (int(testinput[0]) * 2) + (int(testinput[1]) * 2) + int(testinput[2])
#     x = x + (testinput[0]) * (testinput[1]) * (testinput[2])

# print(x)

# print(y)

# def MainFunction(l,w,h):
    # a = l*w
    # b = w*h
    # c = h*l
    # return (2 * a) + (2 * b) + (2 * c) + min(a,b,c)

