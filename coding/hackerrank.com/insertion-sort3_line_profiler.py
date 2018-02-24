#!/usr/bin/env python3
# https://www.quora.com/How-can-I-efficiently-compute-the-number-of-swaps-required-by-slow-sorting-methods-like-insertion-sort-and-bubble-sort-to-sort-a-given-array


# @profile
def merge(left, right):

    tmp = []

    append = tmp.append # optimization!!!
    i, j, inversions = 0, 0, 0
    len_left = len(left)
    len_right = len(right)

    while i < len_left and j < len_right:
        if left[i] <= right[j]:
            append(left[i])
            i += 1
            inversions += j
        else:
            append(right[j])
            j += 1
    inversions += j * (len(left) - i)
    tmp += left[i:]
    tmp += right[j:]
    return inversions, tmp


def merge_sort(arr):
    l = len(arr)
    if l <= 1:
        return 0, arr
    mid = l // 2
    left_inversions, left = merge_sort(arr[:mid])
    right_inversions, right = merge_sort(arr[mid:])
    merge_inversions, merged = merge(left, right)
    inversions = left_inversions + right_inversions + merge_inversions
    return inversions, merged


def insertionSort(arr):
    inv, _ = merge_sort(arr)
    return inv


if __name__ == "__main__":
    fh = open("insert_sort_count.txt")
    num_of_query = int(fh.readline())
    for i in range(num_of_query):
        array_size = fh.readline()
        arr = list(map(int, fh.readline().split(" ")))
        result = insertionSort(arr)
        print(result)


# Wrote profile results to insertion-sort3_line_profiler.py.lprof
# Timer unit: 1e-06 s
#
# Total time: 79.4328 s
# File: ./insertion-sort3_line_profiler.py
# Function: merge at line 5
#
# Line #      Hits         Time  Per Hit   % Time  Line Contents
# ==============================================================
#      5                                           @profile
#      6                                           def merge(left, right):
#      7
#      8   1499985     844382.0      0.6      1.1      tmp = []
#      9
#     10   1499985     887622.0      0.6      1.1      append = tmp.append # optimization!!!
#     11   1499985     852192.0      0.6      1.1      i, j, inversions = 0, 0, 0
#     12   1499985     945704.0      0.6      1.2      len_left = len(left)
#     13   1499985     899859.0      0.6      1.1      len_right = len(right)
#     14
#     15  24545684   15893135.0      0.6     20.0      while i < len_left and j < len_right:
#     16  23045699   17814048.0      0.8     22.4          if left[i] <= right[j]:
#     17  11398432    7852439.0      0.7      9.9              append(left[i])
#     18  11398432    6858440.0      0.6      8.6              i += 1
#     19  11398432    6853327.0      0.6      8.6              inversions += j
#     20                                                   else:
#     21  11647267    8158946.0      0.7     10.3              append(right[j])
#     22  11647267    7011027.0      0.6      8.8              j += 1
#     23   1499985    1302524.0      0.9      1.6      inversions += j * (len(left) - i)
#     24   1499985    1227569.0      0.8      1.5      tmp += left[i:]
#     25   1499985    1195668.0      0.8      1.5      tmp += right[j:]
#     26   1499985     835922.0      0.6      1.1      return inversions, tmp
#
