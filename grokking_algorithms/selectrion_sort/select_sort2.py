#!/usr/bin/env python3
import unittest

def selection_sort(arr):
    for i in range(len(arr)):
        min_pos = i
        for y in range(i + 1, len(arr)):
            if arr[y] < arr[min_pos]:
                min_pos = y
        arr[i], arr[min_pos] = arr[min_pos], arr[i]
    return arr

class TestSelectionSort(unittest.TestCase):

    def setUp(self):
        self.lst = [100, 10, 1000, 1]
        self.lst_sorted = [1, 10, 100, 1000]

    def testResult(self):
        self.assertEqual(selection_sort(self.lst), self.lst_sorted, 'Problem with sorting')

    def testSortedResult(self):
        self.assertEqual(selection_sort(self.lst_sorted), self.lst_sorted, 'Problem with sorting')

if __name__ == '__main__':
    a = [1, 3, 2, -1, 4, 5]
    unittest.main()
