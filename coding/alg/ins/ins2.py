#!/usr/bin/env python

def ins(k):
    for i in range(1, len(k)):
        j = i
        print("j = {}".format(j))
        while j > 0 and k[j-1] > k[j]:
            print(j)
            k[j-1], k[j] = k[j], k[j-1]
            j -= 1
    return k

k = [2,3,4,5,6,7,1]
print(ins(k))
