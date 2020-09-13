#!/usr/bin/env python3

class Solution:
    def reverseBits(self, input):
        '''
        :type input: int
        :rtype: int
        '''

        output = 0
        while input != 0:
            output = output << 1
            if input & 1:
                output = output | 1
            input = input >> 1
        return output

res = Solution()
n = 10
print(f"{n:08b}")
n = res.reverseBits(10)
print(f"{n:08b}")



