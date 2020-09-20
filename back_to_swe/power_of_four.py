#!/usr/bin/env python3

class Solution:
    def powerOfFour(self, input):
        '''
        :type input: int
        :rtype: bool
        '''

        alternatingOddBitMask = 0x55555555
        not_negative = input > 0

        power_of_two = input & (input - 1) == 0

        return bool(not_negative and power_of_two and alternatingOddBitMask)


obj = Solution()
res = obj.powerOfFour(16)
assert res == True

assert obj.powerOfFour(0) == False