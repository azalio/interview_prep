#!/usr/bin/env python3

def bi_search(array, num):
  arr_len = len(array)
  if arr_len == 0:
    return None
  low = 0
  high = arr_len - 1
  while high >= low:
    mid = (low + high)//2
    guess = array[mid]
    if guess == num:
      return mid
    elif guess > num:
      high = mid - 1
    else:
      low = mid + 1
  return None


if __name__ == '__main__':
  array1 = [-1, 0, 1, 2, 3, 4]
  assert bi_search(array1, -1) == 0
  assert bi_search(array1, -10) == None
  assert bi_search(array1, 4) == 5
  
