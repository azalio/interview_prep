#!/usr/bin/env python3

from time import time

def min_distance(str1, str2):
    """str1 -> str2"""
    dp = {}
    # base case
    str1_len = len(str1)
    str2_len = len(str2)
    for i in range(str1_len):
        dp[i,0] = i
    for i in range(str2_len):
        dp[0,i] = i
    # print(f"base dp is: {dp}")

    for j in range(1, str2_len):
        for i in range(1, str1_len):
            # print(f"j: {j}, i:{i}")
            if str1[i] == str2[j]:
                dp[i, j] = dp[i-1, j-1]
            else:
                dp[i, j] = min(dp[i, j-1], dp[i-1, j], dp[i-1, j-1]) + 1 
    return dp[str1_len - 1, str2_len - 2]

start = time()
result = min_distance('intention', 'execution')
print(f"{result}")
end = time()
print(f"time is: {end - start}")