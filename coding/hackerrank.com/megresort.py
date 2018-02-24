#!/usr/bin/env python3


def merge(left, right):
    tmp = []
    i, j = 0, 0
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            tmp.append(left[i])
            i += 1
        else:
            tmp.append(right[j])
            j += 1
    tmp += left[i:]
    tmp += right[j:]
    return tmp


def merge_sort(arr):
    if len(arr) <= 1:
        return arr
    else:
        mid = len(arr)//2
        left = merge_sort(arr[:mid])
        right = merge_sort(arr[mid:])
        return merge(left, right)


arr = [2]
result = merge_sort(arr)
print(result)
