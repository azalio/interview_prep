#!/usr/bin/env python3

def find_smallest(arr):
    smallest_elem = arr[0]
    smallest_pos = 0
    for i in range(1, len(arr)):
        if arr[i] < smallest_elem:
            smallest_elem = arr[i]
            smallest_pos = i
    return smallest_pos

def selection_sort(arr):
    sorted_arr = []
    for i in range(len(arr)):
        pos = find_smallest(arr)
        sorted_arr.append(arr.pop(pos))
    return sorted_arr


if __name__ == '__main__':
    a = [1, 3, 2, -1, 4, 5]
    assert selection_sort(a) == [-1, 1, 2, 3, 4, 5]
    print(a) # -> [] !!!
