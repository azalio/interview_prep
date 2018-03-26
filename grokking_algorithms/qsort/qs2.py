#!/usr/bin/env python3
import random
import unittest


def qsort(lst):
    if len(lst) < 2: return lst
    pivot = random.choice(lst)
    head = qsort([x for x in lst if x < pivot])
    tail = qsort([x for x in lst if x > pivot])
    return head + [x for x in lst if x == pivot] + tail


class TestSort(unittest.TestCase):
    def setUp(self):
        self.lst = [-1, -10, 2, 30, 2]

    def test_sort(self):
        self.assertEqual(qsort(self.lst), sorted(self.lst), 'Problem with qsort')


if __name__ == '__main__':
    unittest.main()
