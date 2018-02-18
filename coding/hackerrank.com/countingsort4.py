#!/usr/bin/env python3
import sys


def count_sort(arr):
    spare_arr = []
    arr_half = len(arr)/2
    for index, value in enumerate(arr):
        if index <= arr_half:
            value[1] = '-'
    for x in range(100):
        spare_arr.append([])
    for i in arr:
        spare_arr[i[0]].append(i[1])
    for index, values in enumerate(spare_arr):
        if len(spare_arr[index]) > 0:
            for y in spare_arr[index]:
                print("{}".format(y), end=' ')
    print("")


#    return(result)

if __name__ == "__main__":
    pass
    # n = int(input().strip())
    # for a0 in range(n):
    #     x, s = input().strip().split(' ')
    #     x, s = [int(x), str(s)]

    arr = [[0,'ab'],[6, 'cd'], [0, 'ef'], [6, 'gh'], [4, 'ij'], [0, 'ab']]
    result = count_sort(arr)
#    print(result)
