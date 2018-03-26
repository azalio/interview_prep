#!/usr/bin/env python3
# https://stackoverflow.com/questions/18262306/quicksort-with-python

import unittest


def qsort(arr):
    if len(arr) < 2:
        return arr
    else:
        pivot = arr[0]
        less = [x for x in arr[1:] if x <= pivot]
        great = [x for x in arr[1:] if x > pivot]
        return qsort(less) + [pivot] + qsort(great)

def qsort2(arr):
    if not arr: return arr
    pivot = arr[0]
    head = qsort2([x for x in arr if x < pivot])
    tail = qsort2([x for x in arr if x > pivot])
    return head + [x for x in arr if x == pivot] + tail


class TestQSort(unittest.TestCase):
    def setUp(self):
        self.arr = [-1, 10, -2, 3, 34]

    def test_qsort(self):
        self.assertEqual(qsort(self.arr), sorted(self.arr), 'Problem with qsort')
        self.assertEqual(qsort2(self.arr), sorted(self.arr), 'Problem with qsort')


if __name__ == '__main__':
    unittest.main()
