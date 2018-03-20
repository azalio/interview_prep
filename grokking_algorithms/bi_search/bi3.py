#!/usr/bin/env python3
import unittest

def bi_search(arr, num):
    len_arr = len(arr)
    low = 0
    high = len_arr - 1
    while high >= low:
        mid = (high + low)//2
        guess = arr[mid]
        if guess == num:
            return mid
        elif guess > num:
            high = mid - 1
        else:
            low = mid + 1
    return None

class TestBiSearch(unittest.TestCase):

    def setUp(self):
        self.lst_sorted = [1, 10, 100, 1000]

    def testResult(self):
        self.assertEqual(bi_search(self.lst_sorted, 10), 1, 'Problem with sorting')
        self.assertEqual(bi_search(self.lst_sorted, 1000), 3, 'Problem with sorting')

    def testSortedResult(self):
        self.assertEqual(bi_search(self.lst_sorted, -1), None, 'Problem with sorting')

if __name__ == '__main__':
    unittest.main()
