#!/usr/bin/env python3

import unittest


def rec_count(arr):
    arr = arr[:]
    if len(arr) == 1:
        return 1
    if len(arr) == 0:
        return 0
    else:
        arr.pop()
        return 1 + rec_count(arr)


class TestCount(unittest.TestCase):
    def setUp(self):
        self.lst = [0, 1, 2, -10, 1000]

    def testCount(self):
        self.assertEqual(rec_count(self.lst),len(self.lst), 'Problem with count')


if __name__ == '__main__':
    unittest.main()
