#!/usr/bin/env python3
from typing import List
import time
# import ipdb

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

def coinChange3(coins: List[int], amount: int):
    dp = [(amount + 1)] * (amount + 1)
    dp[0] = 0
    for i, _ in enumerate(range(len(dp))):
        for coin in coins:
            if (i - coin < 0): continue
            # ipdb.set_trace()
            dp[i] = min(dp[i], 1 + dp[i - coin])
    
    return -1 if dp[amount] == amount + 1 else dp[amount]

start = time.time()
result = coinChange([1,2,5], 20)
end = time.time()
print(end - start)
assert result == 4


start = time.time()
result = coinChange2([1,2,5], 20)
end = time.time()
print(end - start)
assert result == 4


start = time.time()
result = coinChange3([1,2,5], 20)
end = time.time()
print(end - start)
assert result == 4

