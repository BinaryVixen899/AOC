#Intcode

# Testdata = [1,9,10,3,2,3,11,0,99,30,40,50]
# Testdata = [1,1,1,4,99,5,6,0,99]
Testdata = [1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,6,1,19,2,19,9,23,1,23,5,27,2,6,27,31,1,31,5,35,1,35,5,39,2,39,6,43,2,43,10,47,1,47,6,51,1,51,6,55,2,55,6,59,1,10,59,63,1,5,63,67,2,10,67,71,1,6,71,75,1,5,75,79,1,10,79,83,2,83,10,87,1,87,9,91,1,91,10,95,2,6,95,99,1,5,99,103,1,103,13,107,1,107,10,111,2,9,111,115,1,115,6,119,2,13,119,123,1,123,6,127,1,5,127,131,2,6,131,135,2,6,135,139,1,139,5,143,1,143,10,147,1,147,2,151,1,151,13,0,99,2,0,14,0]
z = 0
zcount = 0
count = 0
x = 0 
position1 =0
position2 =0
y = 0

# So the issue here is that it doesn't take 
# into account the four step


for i in Testdata:
    
    if i == 1 and zcount == count:
        x = Testdata[z+1] 
        y = Testdata[z+2] 
        print(x, 'x')
        print(y,'y')
        print(Testdata[x])
        print('test')
        print(Testdata[y])
        sum2 = Testdata[x] + Testdata[y] 
        print(sum2, 'sum2')
        position1 = Testdata[z+3] 
        Testdata[position1] = sum2
        print(Testdata[position1], 'final')
        print(Testdata)
        print(Testdata[z+3], 'answer')
        z+=4
        zcount+=4
        print(z, 'here')
        count+=1
        continue

    elif i == 2 and zcount == count:
       x = Testdata[z+1] 
       print(x,'test')
       y = Testdata[z+2]
       sum1 = Testdata[x] * Testdata[y]
       print(Testdata[z+3])
       position2 = Testdata[z+3] 
       Testdata[position2] = sum1
       z+=4
       zcount+=4
       count+=1
       continue

    elif i == 99 and zcount == count:
       z+=4
       zcount+=4
       count+=1
       continue
    
    else:
       count+=1
       continue

    
       

print(Testdata)

#     pass
# #Process phase 
# count = 0
# combinedlist = []
# for i in (Testdata):
#     newlist = Testdata[0:4]
#     combinedlist.append(newlist)
#     print(newlist)
#     del Testdata[0:4]

# len(combinedlist)

# for x in combinedlist:
#     currentlist = combinedlist[z]
#     z +=1   
#     nextlist = combinedlist[z]
#     currentopcode = currentlist[1]
#     print(str(currentlist) + 'mark')
#     print(str(nextlist) + 'mark2')

# print(currentlist)
# def listmechanism (parameter_list):
   

# # combinedlist[1]
# # print("marker")


# # print(combinedlist)

# # for values in Testdata: 
# #     sliceofthepie = range(0:3)
# #     sliced = slice(3)
# #     i+= 1
# #     print(Testdata[sliced + 'i'])
# # chunkIt(Testdata, 3)
# # def chunkIt(seq, num):
# #     avg = len(seq) / float(num)
# #     out = []
# #     last = 0.0

# #     while last < len(seq):
# #         out.append(seq[int(last):int(last + avg)])
# #         last += avg

# #     return out

# # exit()


# # step forward 4 



# # 1,0,0,0,99 becomes 2,0,0,0,99 (1 + 1 = 2).
# # 2,3,0,3,99 becomes 2,3,0,6,99 (3 * 2 = 6).
# # 2,4,4,5,99,0 becomes 2,4,4,5,99,9801 (99 * 99 = 9801).
# # 1,1,1,4,99,5,6,0,99 becomes 30,1,1,4,2,5,6,0,99.
# # Once you have a working computer, the first step is to restore the gravity assist program (your puzzle input) to the "1202 program alarm" state it had just before the last computer caught fire. To do this, before running the program, replace position 1 with the value 12 and replace position 2 with the value 2. What value is left at position 0 after the program halts?