#!/usr/bin/env python3

import unittest


class TestSort(unittest.TestCase):

    def setUp(self):
        self.lst = [3, -1, 4, 1, 9]

    def test_sort(self):
        self.assertEqual(select_sort(self.lst), sorted(self.lst), 'Problem with sorting')


def select_sort(lst):
    for i in range(len(lst)):
        min_pos = i
        for y in range(i+1, len(lst)):
            if lst[y] < lst[min_pos]:
                min_pos = y
        lst[min_pos], lst[i] = lst[i], lst[min_pos]
    return lst


if __name__ == '__main__':
    unittest.main()
