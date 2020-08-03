#!/usr/bin/env python3

class Solution:
  def twoSum(self, nums, target):
        seen = {}
        for i, v in enumerate(nums):
            remaining = target - v
            if remaining in seen:
                return [seen[remaining], i]
            seen[v] = i
        return []
       
ins = Solution()

nums = [3,3]
target = 6

print(ins.twoSum(nums,target)) 
