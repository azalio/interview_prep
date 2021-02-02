#!/usr/bin/env python3


class Solution:
    def countSubarrays(self, arr, k):  # O(n^2)
        '''
        :type arr: list of int
        :type k: int
        :rtype: int
        '''
        sum_of_nums = []  # O(1)
        sum = 0  # O(1)
        for i in arr:  # O(n)
            sum = sum + i  # O(1)
            sum_of_nums.append(sum)  # O(1)

        i = len(arr) - 1  # O(1)

        found = 0  # O(1)

        while i >= 0:  # O(n^2)
            if sum_of_nums[i] == k:  # O(1)
                # print(arr[:i+1])
                found += 1  # O(1)
            j = i - 1
            while j >= 0:  # O(n)
                if sum_of_nums[i] - sum_of_nums[j] == 0:  # O(1)
                    # print(arr[j+1:i+1])
                    found += 1  # O(1)
                j -= 1  # O(1)
            i -= 1  # O(1)
        return found  # O(1)


obj = Solution()
arr = [1, 0, -1, 1]
k = 0
assert obj.countSubarrays(arr, k) == 4

obj = Solution()
arr = [3, 7, -4, -2, 1, 5]
k = 3
assert obj.countSubarrays(arr, k) == 2
