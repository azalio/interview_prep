#!/usr/bin/env python

def bi_search(arr, num):
    arr_len = len(arr)
    low = 0
    high = arr_len - 1
    while high >= low:
        mid = (high + low)//2
        guess = arr[mid]
        if guess == num:
            return mid
        elif guess > num:
            high = mid - 1
        else:
            low = mid + 1
    return None


if __name__ == '__main__':
    arr = [-1, 0, 1, 2, 3, 4]
    assert bi_search(arr, -1) == 0
    assert bi_search(arr, 4) == 5
    assert bi_search(arr, -2) == None
