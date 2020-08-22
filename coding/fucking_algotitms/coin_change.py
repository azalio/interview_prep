#!/usr/bin/env python3
from typing import List
import time

def coinChange(coins: List[int], amount: int):
    def dp(n):
        # base case
        if n == 0: return 0
        if n < 0: return -1
        # to minimize it is to initialize it to infinity
        res = float('INF')
        for coin in coins:
            subproblem = dp(n - coin)
            # No solution to subproblem, skip
            if subproblem == -1: continue
            res = min(res, 1 + subproblem)

        return res if res != float('INF') else -1

    return dp(amount)


def coinChange2(coins: List[int], amount: int):
    # memo
    memo = dict()
    def dp(n):
        # Check the memo to avoid double counting
        if n in memo: return memo[n]

        if n == 0: return 0
        if n < 0: return -1
        res = float('INF')
        for coin in coins:
            subproblem = dp(n - coin)
            if subproblem == -1: continue
            res = min(res, 1 + subproblem)

        # note on memo
        memo[n] = res if res != float('INF') else -1
        return memo[n]

    return dp(amount)



start = time.time()
result = coinChange([1,2,5], 30)
end = time.time()
print(end - start)

start = time.time()
result = coinChange2([1,2.5], 30)
end = time.time()
print(end - start)