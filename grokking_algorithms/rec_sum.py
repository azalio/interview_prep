#!/usr/bin/env python3
import unittest


class testSum(unittest.TestCase):

    def setUp(self):
        self.lst = [-1, 0, 2, 1, 10]

    def testSum(self):
        self.assertEqual(summ(self.lst), sum(self.lst), 'Problem with sum')

def summ(arr):
    arr = arr[:]
    if len(arr) == 1:
        return arr[0]
    elif len(arr) == 0:
        return 0
    else:
        return arr.pop() + summ(arr)

if __name__ == '__main__':
    unittest.main()
