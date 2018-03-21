#!/usr/bin/env python3
import unittest

class TestBiSearch(unittest.TestCase):
    
    def setUp(self):
        self.arr = [1, 2, 3, 4, 5]

    def testResult(self):
        self.assertEqual(bi_search(self.arr, 1), 0, 'Problem with search')
        self.assertEqual(bi_search(self.arr, 5), 4, 'Problem with search')

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

if __name__ == "__main__":
    unittest.main()
