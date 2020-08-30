# --- Day 4: The Ideal Stocking Stuffer ---
# Santa needs help mining some AdventCoins (very similar to bitcoins) to use as gifts for all the economically forward-thinking little girls and boys.

# To do this, he needs to find MD5 hashes which, in hexadecimal, start with at least five zeroes. The input to the MD5 hash is some secret key (your puzzle input, given below) followed by a number in decimal. To mine AdventCoins, you must find Santa the lowest positive number (no leading zeroes: 1, 2, 3, ...) that produces such a hash.

# For example:

# If your secret key is abcdef, the answer is 609043, because the MD5 hash of abcdef609043 starts with five zeroes (000001dbbfa...), and it is the lowest such number to do so.
# If your secret key is pqrstuv, the lowest number it combines with to make an MD5 hash starting with five zeroes is 1048970; that is, the MD5 hash of pqrstuv1048970 looks like 000006136ef....

#MD5 hash start with at least five zeros
#MD5Hash = Lowest Positive Number * Secret Key 

# A better way to understand this would be to actually calculate at a lower level how stuff works in terms of registers
# However I am tired and pressed for time, so I am likely going to do it the unceremonious way and just calculate it out 
# This will howerver probably come at a cost for performance
#Although it can still be optimize
#Ultimately what we're doing here is figuring out the othe rhalf of a hash 

#This is really inelegant
#But hey it works 
# --- Part Two ---
# Now find one that starts with six zeroes.
#Just change it between 5 zeros 6 zeros for part 2
import hashlib
#So the key is making sure that this isn't one
# mykey = bytes(mykey)
for i in range(1,100000000000000):
    # print(i)
    test = 'bgvyzdsv' + str(i) 
    # print(test)
    test = bytes(test,encoding='utf8') 
    result = hashlib.md5(test)
    result = result.hexdigest()
    # print(result)
    if '000000' == (result[0:6]):
        print(result)
        print(i)
        quit()
    # print(lastresult)
#have this throw it out when we ge tthe lowest 



# use == instead of is not
    # print()hexdigest

# result = hashlib.md5('abcdef' + mykey) #Okay so what I've proven here is that I can do the concatentaiton
# print(result.hexdigest())                   
#There are potential ways to make this easier, and would be better if we knew precisely how MD5 hashes worked but nonetheless....
#If the answer doesn't start with 0 we can throw it out
#so basically we're going to be eliminating large swaths of combinations

# result = hashlib.mdr(b'abcdef') + 100000:999999
# for i in result if result[0:5] != 0 
#See this is how 
# An MD5 hash is created by taking a string of an any length and encoding it into a 128-bit fingerprint. Encoding the same string using the MD5 algorithm will always result in the same 128-bit hash output.