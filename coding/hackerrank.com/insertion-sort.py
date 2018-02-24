#!/usr/bin/env python3


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


def insertionSort(A):
    # Complete this function
    res = 0
    bit = Bit(max(A)+1)
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
