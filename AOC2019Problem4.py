# --- Day 4: Secure Container ---
# You arrive at the Venus fuel depot only to discover it's protected by a password. The Elves had written the password on a sticky note, but someone threw it out.

# However, they do remember a few key facts about the password:

# It is a six-digit number.
# The value is within the range given in your puzzle input.
# Two adjacent digits are the same (like 22 in 122345).
# Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
# Other than the range rule, the following are true:

# 111111 meets these criteria (double 11, never decreases).
# 223450 does not meet these criteria (decreasing pair of digits 50).
# 123789 does not meet these criteria (no double).
# How many different passwords within the range given in your puzzle input meet these criteria?

# Your puzzle input is 367479-893698.

# inputvalue = '367479-893698'

# 111111 meets these criteria (double 11, never decreases).
# 223450 does not meet these criteria (decreasing pair of digits 50).
# 123789 does not meet these criteria (no double).
from collections import Counter
success = 0
success2 =0



def Calculate(data):
    global success
    lastentry='null'
    bothconditionspassed = False

    for entry in data[0:len(data)]:
        # print('entry \n' + str(entry))
        # print('lastentry \n' + str(lastentry))
        if lastentry == 'null': 
            lastentry = entry
            continue
        else:
            if entry >= lastentry:
                if lastentry == entry: 
                    bothconditionspassed = True
                    lastentry = entry
                    continue
                else:
                    lastentry = entry
                    continue
            else:
                # print('failure')
                bothconditionspassed = False
                break

    if bothconditionspassed == True:
        # print("sucess!")
        success = success + 1
        # print('True')

            
def digits(n):
   # None, not corect for n <= 0
   ret = []
   while n > 0:
        ret.append(n % 10)
        n = n // 10
   ret.reverse()
   return ret


def Calculate2(data):
    global success2
    lastentry='null'
    bothconditionspassed = False
    

    for entry in data[0:len(data)]:

        
                
            #need something like that 
    

        # print('entry \n' + str(entry))
        # print('lastentry \n' + str(lastentry))
        if lastentry == 'null': 
            lastentry = entry
            continue
        else:
            if entry >= lastentry:
                if lastentry == entry: 
                    bothconditionspassed = True
                    lastentry = entry
                    continue
                else:
                    lastentry = entry
                    continue
            else:
                # print('failure')
                bothconditionspassed = False
                break

    if bothconditionspassed == True:
        
        # print("sucess!")
        success2 = success2 + 1
        # print('True')

datalist = [digits(k) for k in range(367479,893698)]


for i in datalist:
    Calculate(i)

print(success)
    

# An Elf just remembered one more important detail: the two adjacent matching digits are not part of a larger group of matching digits.

# Given this additional criterion, but still ignoring the range rule, the following are now true:

# 112233 meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
# 123444 no longer meets the criteria (the repeated 44 is part of a larger group of 444).
# 111122 meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).
# How many different passwords within the range given in your puzzle input meet all of the criteria?




# for i in datalist2:
#     Calculate2(i)

print(success)  
a = 0
testdict ={}
data = [digits(k) for k in range(367479,893698)]     
datalist2 = [digits(k) for k in range(100,110)]
for i in datalist2:
    for a in [i]:
        z = map(lambda x: x *2, a)
        print(list(z))



print(testdict)

# for x in datadicttest:
#     print(datadicttest)

# print(datalist2)

# for x in datadict.values():
#             print(datadict)
#             print(x)
#             if x % 2 == 0:
#                 print("okay")
#             else:
#                 break