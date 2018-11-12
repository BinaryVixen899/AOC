# A new system policy has been put in place that requires all accounts to use a passphrase instead of simply a password. 
# A passphrase consists of a series of words (lowercase letters) separated by spaces.

# To ensure security, a valid passphrase must contain no duplicate words.

#could easily do the things from before in the, i think 2nd problem

#Initialization
passphrase = "aa bb cc dd ee"
currentitem = ""
currentitemstep = 0 

#deciphering
passphraselist = list(passphrase.split(" "))
print(passphraselist)

listlength = len(passphraselist) 

for  phrase in passphraselist: 
    currentitem = phrase
    
    for phrase in passphraselist[currentitemstep + 1:]:
        if phrase == currentitem:
            print("INVALID")
            break
    
    currentitemstep += 1


##


# For example:

# aa bb cc dd ee is valid.
# aa bb cc dd aa is not valid - the word aa appears more than once.
# aa bb cc dd aaa is valid - aa and aaa count as different words.
# The system's full passphrase list is available as your puzzle input. How many passphrases are valid?