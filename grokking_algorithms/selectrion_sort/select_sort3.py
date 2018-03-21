#!/usr/bin/env python3

import unittest


def select_sort(arr):
    for i in range(len(arr)):
        min_pos = i
        for y in range(i + 1, len(arr)):
            if arr[y] < arr[min_pos]:
                min_pos = y
        arr[i], arr[min_pos] = arr[min_pos], arr[i]
    return arr


class testSelectSort(unittest.TestCase):

    def setUp(self):
        self.arr = [1, -1, 2, 0, -10]

    def testSort(self):
        self.assertEqual(select_sort(self.arr), [-10, -1, 0, 1, 2], 'Problem with sort')


if __name__ == '__main__':
    unittest.main()
