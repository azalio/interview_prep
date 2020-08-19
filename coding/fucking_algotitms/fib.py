#!/usr/bin/env python3
import time
import sys
sys.setrecursionlimit(10000)

fib_dict = {}

def fibonachi(num):
    # print(num)
    if num == 1 or num == 2:
        return 1
    if fib_dict.get(num, -1) != -1:
        return fib_dict[num]
    fib_dict[num] = fibonachi(num - 1) + fibonachi(num - 2)
    return fib_dict[num]


def fib2(num):
    if num <= 2: return 1
    curr, prev = 2,1
    for _ in range(3,num):
        sum = prev + curr
        prev = curr
        curr = sum
    return sum
    

start = time.time()
_ = fibonachi(9999)
end = time.time()
print(end - start)

start = time.time()
_ = fib2(9999)
end = time.time()
print(end - start)