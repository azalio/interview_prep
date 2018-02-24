#!/usr/bin/env python3


def find_closest(aList):
    min = aList[1] - aList[0]
    position = [0]
    for i in range(2, len(aList)):
        min_new = aList[i] - aList[i-1]
        if min == min_new:
            position.append(i)
        elif min > min_new:
            min = min_new
            position = [i]
    return position

def insert_sort(aList):
    for i in range(1, len(aList)):
        tmp = aList[i]
        k = i
        while k > 0 and tmp < aList[k - 1]:
            aList[k] = aList[k - 1]
            k -= 1
        aList[k] = tmp
    positions = find_closest(aList)
    for i in positions:
        print("{} {}".format(aList[i-1], aList[i]), end = " ")





if __name__ == "__main__":
    n = 10
    arr = [-20, -3916237, -357920, -3620601, 7374819, -7330761, 30, 6246457, -6461594, 266854]
    result = insert_sort(arr)
    # positions = find_closest(result)
    # for i in positions:
    #     print("{} {}".format(result[i-1], result[i]), end = " ")
    # print("")

