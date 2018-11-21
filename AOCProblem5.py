# --- Day 5: A Maze of Twisty Trampolines, All Alike ---
# The message includes a list of the offsets for each jump. Jumps are relative: 
# -1 moves to the previous instruction, and 2 skips the next one. 
# Start at the first instruction in the list. The goal is to follow the jumps until one leads outside the list.

# In addition, these instructions are a little strange; after each jump, the offset of that instruction increases by 1.
# So, if you come across an offset of 3, you would move three instructions forward,
# but change it to a 4 for the next time it is encountered.

# For example, consider the following list of jump offsets:
# 0
# 3
# 0
# 1
# -3
# Positive jumps ("forward") move downward; negative jumps move upward. 
# For legibility in this example, these offset values will be written all on one line, 
# with the current instruction marked in parentheses. The following steps would be taken before an exit is found:

# (0) 3  0  1  -3  - before we have taken any steps.
# (1) 3  0  1  -3  - jump with offset 0 (that is, don't jump at all). Fortunately, the instruction is then incremented to 1.
#  2 (3) 0  1  -3  - step forward because of the instruction we just modified. The first instruction is incremented again, now to 2.
#  2  4  0  1 (-3) - jump all the way to the end; leave a 4 behind.
#  2 (4) 0  1  -2  - go back to where we just were; increment -3 to -2.
#  2  5  0  1  -2  - jump 4 steps forward, escaping the maze.
# In this example, the exit is reached in 5 steps.

# How many steps does it take to reach the exit?

#---Variables---#
#Okay I need to really think about how I'm defining things here
values = [0,3,0,1,-3]

#initialization
currentposition = 0
coordinatevalue = values[currentposition]
step = 0
coordinatelength = len(values)

#---Programs---#
#need a coordinate increment AND a value increment, heck and a move one too but the move should go down in coordinate move 
# def CoordinateIncrement (value):
#     coordinatevalue = value + 1
#     return coordinatevalue


#Coordinate Calculation

#-----Coordinate Move-----#
print("start")
print("coordinatevalue",coordinatevalue)
print("currentposition",currentposition)

while currentposition <=4:
    print(values[currentposition])
    currentposition = currentposition + values[currentposition]
    
    #It should attempt to move first
    
    if currentposition <=4:
        #Increment
        values[currentposition] = values[currentposition] + 1
    else:
        print("Break")
        break
    
    #Position Move
    
    
    step += 1
    print("step",step)
    
    print("coordinatevalue",coordinatevalue)
    print("currentposition",currentposition)
    

#okay i am doing something wrong here
    
#<------------------ #here we go, coordinvate value is not changing the actual value of the list, which is the problem



