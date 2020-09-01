# --- Day 5: Doesn't He Have Intern-Elves For This? ---
# Santa needs help figuring out which strings in his text file are naughty or nice.

# A nice string is one with all of the following properties:

# It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
# It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
# It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.
# For example:

# ugknbfddgicrmopn is nice because it has at least three vowels (u...i...o...), a double letter (...dd...), and none of the disallowed substrings.
# aaa is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.
# jchzalrnumimnmhp is naughty because it has no double letter.
# haegwjzuvuyypxyu is naughty because it contains the string xy.
# dvszwmarrgswjxmb is naughty because it contains only one vowel.
# How many strings are nice?

#so what if we just chunk this 
#And do three operations for each

import collections
import re
#vowels
vowels = 'aeiou'
# duplicates
bad = ['ab','cd','pq','xy']

puzzleinput = 'ugknbfddgicrmopn'
testlist = []

 

#Need vowels 
#Need to put it all together 

#Vowel count
c = collections.Counter(puzzleinput)
result = "".join(set(puzzleinput))
for i in result:  #get unique value
    print(f"letter {i} lettercount {c[i]}")
#Perhaps delete key? if i not in 


 
 
#Got a way to check for duplicates
testlist = (re.findall(r'..', puzzleinput))
for i in testlist:
    print(i)
    if i[0] == i[1]:
        print("yay")
        hasduplicates = True
        break

for i in bad: #Okay so got a wayt to find the bad words
    if i in puzzleinput:
        print("hahano")
        quit()
    else:
        print("yay")

#Okay 
# print(testlist)