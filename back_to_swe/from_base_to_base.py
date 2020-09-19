#!/usr/bin/env python3

import string

base_list = list(range(10)) # O(n)
base_list.extend(list(string.ascii_uppercase)) # O(n)

base_dict = dict() O(1)

for i, v in enumerate(base_list): # O(n)
    base_dict[str(v)] = i

# print(f"base_list: {base_list}")


class Solution: # O(n)
    def changeBase(self, numAsString: str, b1: int, b2: int) -> str:
        '''
        :type numAsString: str
        :type b1: int
        :type b2: int
        :rtype: str
        '''
        if b1 == b2: # O(1)
            return numAsString

        if b1 == 10:
            return self.change_from_10_to_b2(numAsString, b2)
        else:
            to_10 = self.change_to_10(numAsString, b1)
            if b2 == 10:
                return to_10
            else:
                return self.change_from_10_to_b2(to_10, b2)

    def change_from_10_to_b2(self, to_10: str, b2: int):
        """
        10 -> 2:
        10/2 -> 10 - 10 -> 0 (5)
        5/2 -> 5 - 4 -> 1 (2)
        2/2 -> 2 - 2 -> 0 (1)
        1010
        """
        to_10 = int(to_10)
        quotient = 37
        res = []
        while quotient >= b2: # O(logb2(to_10)) -> O(logb(n))
            quotient = int(to_10/b2)
            num = to_10 - quotient*b2
            res.append(str(base_list[num]))
            to_10 = quotient
            # print(f"quotient: {quotient}")
        res.append(str(base_list[quotient]))

        # print(f"res: {res}")
        res.reverse()
        return ''.join(res)

    def change_to_10(self, numAsString: str, b1: int):
        """
        1010 = 1*2**3 + 0*2**2 + 1*2**1 + 0*2**0
        """
        to_10 = 0
        for i, v in enumerate(numAsString): # O(n)
            res = base_dict[v] * b1**(len(numAsString) -1 - i)
            # print(f"res: {res}")
            to_10 += res
        return str(to_10)


obj = Solution()
assert obj.changeBase("123", 4, 10) == "27"
assert obj.changeBase("123", 10, 10) == "123"
assert obj.changeBase("10", 10, 2) == "1010"
assert obj.changeBase("12", 10, 2) == "1100"
assert obj.changeBase("123", 4, 16) == "1B"