#!/usr/bin/env python

import sys

# def solve(a0, a1, a2, b0, b1, b2):
def solve(*argv):
    # Complete this function
    a, b = 0, 0
    for i in range(0,3):
        if argv[i] > argv[i+3]:
            a += 1
        elif argv[i] < argv[i+3]:
            b += 1
    return [a, b]


# a0, a1, a2 = raw_input().strip().split(' ')
# a0, a1, a2 = [int(a0), int(a1), int(a2)]
# b0, b1, b2 = raw_input().strip().split(' ')
b0, b1, b2 = 3, 6, 10
# b0, b1, b2 = [int(b0), int(b1), int(b2)]
result = solve(a0, a1, a2, b0, b1, b2)
print " ".join(map(str, result))



