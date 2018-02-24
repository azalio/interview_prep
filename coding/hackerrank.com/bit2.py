#!/usr/bin/env python3
# https://www.quora.com/How-can-I-efficiently-compute-the-number-of-swaps-required-by-slow-sorting-methods-like-insertion-sort-and-bubble-sort-to-sort-a-given-array

class Bit:
    def __init__(self, n):
        sz = 1
        while n > sz:
            sz *= 2
        self.size =sz
        self.data = [0]*sz

    def sum(self, i):
        s = 0
        while i > 0:
            s += self.data[i]
            i -= i & -i
        return s

    def add(self, i, x):
        while i < self.size:
            self.data[i] += x
            i += i & -i


def inversions(A):
    res = 0
    bit = Bit(max(A)+1)
    for i, v in enumerate(A):
        # print("i: {}, v: {}".format(i, v))
        res += i - bit.sum(v)
        bit.add(v, 1)
        print(bit.data)
    return res


if __name__ == '__main__':
    arr = [1, 2, 4, 3]
    print(inversions(arr))
