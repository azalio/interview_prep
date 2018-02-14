#!/usr/bin/env python3
import sys

def insertionSort1(n, arr):
    # Complete this function
    n = arr[-1]
    for i in range(len(arr)-1, 0, -1):
        if arr[i-1] > n:
            arr[i] = arr[i-1]
        elif arr[i-1] == n:
            arr[i] = n
            break
        elif arr[i-1] < n:
            arr[i] = n
            break
        print(" ".join(map(str, arr)))
    else:
        arr[0] = n
    print(" ".join(map(str, arr)))

if __name__ == "__main__":
#    n = int(input().strip())
#    arr = list(map(int, input().strip().split(' ')))
    n = 5
    arr = [2,3,4,5,6,7,8,9,10,1]
    insertionSort1(n, arr)
