#!/usr/bin/env python3
import math

class Solution:
    def isPalindrome(self, x: int) -> bool: # O(n)
        '''
        :type x: int
        :rtype: bool
        '''
        if x < 0: # O(1)
            return False # O(1)
        if x < 10: # O(1)
            return True # O(1)

        digits = math.floor(math.log(x, 10)) + 1 # O(1)
        mask = 10**(digits - 1) # O(1)
        # print(f"digits: {digits}")
        for _ in range(digits//2): # O(n)
            begin = x//mask # O(1)
            end = x%10 # O(1)
            if begin != end: # O(1)
                return False # O(1)
            x = x%mask # O(1)
            x = x//10 # O(1)
            mask = mask//10 # O(1)
        return True # O(1)


obj = Solution()
assert obj.isPalindrome(121), True
assert obj.isPalindrome(212), False