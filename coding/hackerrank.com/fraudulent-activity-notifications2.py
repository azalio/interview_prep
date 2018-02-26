#!/usr/bin/env python3


def merge(left, right):
    tmp = []
    append = tmp.append
    i, j = 0, 0
    len_left = len(left)
    len_right = len(right)
    while i < len_left and j < len_right:
        if left[i] <= right[j]:
            append(left[i])
            i += 1
        else:
            append(right[j])
            j += 1
    tmp += left[i:]
    tmp += right[j:]
    return tmp


def merge_sort(arr):
    arr_len = len(arr)
    if arr_len <= 1:
        return arr
    mid = arr_len//2
    left = merge_sort(arr[:mid])
    right = merge_sort(arr[mid:])
    return merge(left, right)


def activityNotifications(expenditure, window):
    # Complete this function
    notify, i = 0, 0
    odd = (window%2)
    # print(expenditure)
    # print(merge_sort(expenditure))
    mid_list = merge_sort(expenditure)
    mid_position = len(mid_list[:window])//2
    expend_len = len(expenditure)

    while i + window < expend_len:
        shift = expenditure[i:window + i]
        # print(shift)
        # mid_list = merge_sort(shift)
        # mid_position = len(mid_list)//2
        if odd:
            mid = mid_list[mid_position]
        else:
            mid = (mid_list[mid_position-1] + mid_list[mid_position])/2
        print("mid: {}".format(mid))
        # print(mid)
        # shift = expenditure[i:window+i]
        # print(shift)
        # mid_position = (len(shift)//2)
        # mid = shift[mid_position]
        # print("mid: {}".format(mid))
        if expenditure[i+window] >= 2*mid:
            notify += 1
        i += 1
        mid_position += 1
    return notify


if __name__ == "__main__":
    with open("fraudulent-activity-notifications.txt") as fh:
        n, d = fh.readline().split(" ")
        n, d = int(n), int(d)
        expenditure = fh.readline().split(" ")
        expenditure = list(map(int, expenditure))
    days = n
    window = d
    expenditure = expenditure[0:11001]

    days = 9
    window = 4
    expenditure = [1, 1, 4, 2, 3, 6, 8, 4, 5]


    notify = activityNotifications(expenditure, window)
    print(notify)