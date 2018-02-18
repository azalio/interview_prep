#!/usr/bin/env python3

def insertion_sort(l):
    for i in range(1, len(l)):
        j = i-1
        key = l[i]
        while (j >= 0) and (l[j] > key):
           l[j+1] = l[j]
           j -= 1
        l[j+1] = key


m = 6
ar = [4,1]
insertion_sort(ar)
print(" ".join(map(str,ar)))

