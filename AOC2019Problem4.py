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
success = 0

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

datalist = [digits(k) for k in range(367479,893698)]


for i in datalist:
    Calculate(i)

print(success)
    









