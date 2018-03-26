#!/usr/bin/env python3

import unittest
import random


class TestQSort(unittest.TestCase):
    def setUp(self):
        self.lst = [-2, -3, 0, 1, 10, -99]

    def test_q_sort(self):
        self.assertEqual(qsort(self.lst), sorted(self.lst), 'Problem with qsort')


def qsort(lst):
    if len(lst) < 2: return lst
    pivot = random.choice(lst)
    head = qsort([x for x in lst if x < pivot])
    tail = qsort([x for x in lst if x > pivot])
    return head + [x for x in lst if x == pivot] + tail


if __name__ == '__main__':
    unittest.main()
