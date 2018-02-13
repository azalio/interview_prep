#!/usr/bin/env python3

import sys

def plusMinus(arr):
    # Complete this function
    a, b, c = 0, 0, 0
    for i in arr:
        if i > 0:
            a += 1
        elif i < 0:
            b += 1
        else:
            c += 1

    print("{0:.6f}".format(a/len(arr)))
    print("{0:.6f}".format(b/len(arr)))
    print("{0:.6f}".format(c/len(arr)))

if __name__ == "__main__":
#    n = int(input().strip())
#    arr = list(map(int, input().strip().split(' ')))
    arr = [2, -2, 0, -1, 1, 0, 1]
    plusMinus(arr)

