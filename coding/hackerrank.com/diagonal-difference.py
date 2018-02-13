#!/usr/bin/env python
import sys

def diagonalDifference(a):
    # Complete this function
    print(a)
    e, b = 0, 0
    for i in range(len(a[0])):
        e += a[i][i]
        b += a[i][-1-i]
    return abs(e - b)

if __name__ == "__main__":
#    n = int(raw_input().strip())
    a = [[11, 2, 4], [4, 5, 6], [10, 8, -12]]
#    for a_i in xrange(n):
#        a_temp = map(int,raw_input().strip().split(' '))
#        a.append(a_temp)
    result = diagonalDifference(a)
    print result

