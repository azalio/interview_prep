#!/usr/bin/env python3

import unittest


class TestSelectSort(unittest.TestCase):
    def setUp(self):
        self.lst = [0, -1, 2, -10, 3, 10]

    def test_select(self):
        self.assertEqual(select_sort(self.lst), sorted(self.lst), 'Problem with select sort')


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
