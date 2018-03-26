#!/usr/bin/env python3

import unittest


def find_max(arr):
    try:
        max_num = arr[0]
    except IndexError:
        return None
    try:
        for i in range(1, len(arr)):
            if arr[i] > max_num:
                max_num = arr[i]
    except IndexError:
        return max_num
    return max_num


class TestMax(unittest.TestCase):
    def setUp(self):
        self.lst = [1, 2, 3, 4]

    def testMax(self):
        self.assertEqual(find_max(self.lst), max(self.lst), 'Problem with max func()')


if __name__ == '__main__':
    unittest.main()
