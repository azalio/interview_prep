#!/usr/bin/env python3
from array import array

class Bit:
    def __init__(self):
        self.size = 10000001
        self.data = array('I', [0]*10000001)

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


def insertionSort(A):
    # Complete this function
    res = 0
    bit = Bit()
    for i, v in enumerate(A):
        res += i - bit.sum(v)
        bit.add(v, 1)
    return res

if __name__ == "__main__":
    fh = open("insert_sort_count.txt")
    num_of_query = int(fh.readline())
    for i in range(num_of_query):
        array_size = fh.readline()
        arr = list(map(int, fh.readline().split(" ")))
        result = insertionSort(arr)
        print(result)
