#!/usr/bin/env python3

class Solution:
  def twoSum(self, nums, target):
    result = [0,0]
    for idx, i in enumerate(nums):
      result[0] = idx
      for idx2, x in enumerate(nums[result[0] + 1:]):
        if i + x == target:
          result[1] = idx2 + result[0] + 1
          return result
            
       
ins = Solution()

nums = [1, 11, 15, 2, 7 ]
target = 9

print(ins.twoSum(nums,target)) 
