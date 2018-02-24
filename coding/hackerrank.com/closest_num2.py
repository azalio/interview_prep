#!/usr/bin/env python3
import sys


def find_closest(aList):
    print("Starting find_closest")
    # aList = insert_sort(aList)
    aList.sort()
    print("Insert sort ended")
    min = aList[1] - aList[0]
    position = [0]
    print("find min")
    for i in range(2, len(aList)):
        min_new = aList[i] - aList[i-1]
        if min == min_new:
            position.append(i)
        elif min > min_new:
            min = min_new
            position = [i]
    print("find min ended")
    numbers = []
    for i in position:
        numbers.append(aList[i-1])
        numbers.append(aList[i])
    print("append ended")
    return numbers


def insert_sort(aList):
    print("Starting insert_sort")
    for i in range(1, len(aList)):
        tmp = aList[i]
        k = i
        while k > 0 and tmp < aList[k - 1]:
            aList[k] = aList[k - 1]
            k -= 1
        aList[k] = tmp
    return aList

if __name__ == "__main__":
    # n = int(input().strip())
    # arr = list(map(int, input().strip().split(' ')))
    with open("closest_num_test_case.txt") as fh:
        line = fh.read()
    print("file readed")
    arr = [int(x) for x in line.split(' ')]
    print("line splitted")
    result = find_closest(arr)
    print (" ".join(map(str, result)))


