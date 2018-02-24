#!/usr/bin/env python3
# https://www.quora.com/How-can-I-efficiently-compute-the-number-of-swaps-required-by-slow-sorting-methods-like-insertion-sort-and-bubble-sort-to-sort-a-given-array


def merge(left, right):
    tmp = []
    i, j, inversions = 0, 0, 0
    len_left = len(left)
    len_right = len(right)

    while i < len_left and j < len_right:
        if left[i] <= right[j]:
            tmp.append(left[i])
            i += 1
            inversions += j
        else:
            tmp.append(right[j])
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
