#!/usr/bin/env python3
import sys

def insertionSort1(arr):
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
    else:
        arr[0] = n
    return arr

def insertionSort2(n, arr):
    # Complete this function
    for i in range(1, len(arr)):
        arr[0:i+1] = insertionSort1(arr[0:i+1])
        print(" ".join(map(str,arr)))


if __name__ == "__main__":
#    n = int(input().strip())
#    arr = list(map(int, input().strip().split(' ')))
    n = 6
    arr = [1,4,3,5,6,2]
    insertionSort2(n, arr)
