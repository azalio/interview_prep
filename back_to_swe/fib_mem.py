#!/usr/bin/env python3

cache = dict()

def fib(n: int) -> int:
    # base case:
    if n == 0:
        return 0
    if n == 1:
        return 1
    num = cache.get(n, None)
    if num:
        return num
    else:
        cache[n] = fib(n - 1) + fib(n - 2)
        return cache[n]

res = fib(900)
print(f"res: {res}")
    
