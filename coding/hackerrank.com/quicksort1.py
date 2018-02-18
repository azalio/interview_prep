#!/usr/bin/env python3
import sys

def quickSort(arr):
    # Complete this function
    left, right, eq = [], [], []
    p = arr[0]
    for i in arr:
        if p < i:
            right.append(i)
        elif p > i:
            left.append(i)
        else:
            eq.append(i)
    return (left+eq+right)

if __name__ == "__main__":
    arr = [4,5,3,7,2]
    result = quickSort(arr)
    print (" ".join(map(str, result)))
