#!/usr/bin/env python3

def ins_sort(k):
    for i in range(1,len(k)):    #since we want to swap an item with previous one, we start from 1
        j = i                    #bcoz reducing i directly will mess our for loop, so we reduce its copy j instead
        temp = k[j]              #temp will be used for comparison with previous items, and sent to the place it belongs
        while j > 0 and temp < k[j-1]: #j>0 bcoz no point going till k[0] since there is no seat available on its left, for temp
            k[j] = k[j-1] #Move the bigger item 1 step right to make room for temp
            j=j-1 #take k[j] all the way left to the place where it has a smaller/no value to its left.
        k[j] = temp
    return k

arr = [3, 2, 4]
print(ins_sort(arr))
