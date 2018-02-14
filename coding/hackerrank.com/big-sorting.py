#!/usr/bin/env python3
import sys

def bigSorting(arr):
    # Complete this function
    return sorted(arr, key=int)

if __name__ == "__main__":
#    n = int(input().strip())
    arr = ['35', '1', '3', '10', '3', '5']
#    arr_i = 0
#    for arr_i in range(n):
#       arr_t = str(input().strip())
#       arr.append(arr_t)
    result = bigSorting(arr)
    print ("\n".join(map(str, result)))
