#!/usr/bin/env python3

def merge_sort(not_sorted_list) -> list:
    if len(not_sorted_list) <= 1:
        return not_sorted_list
    
    diviner = len(not_sorted_list)//2
    left = merge_sort(not_sorted_list[:diviner])
    right = merge_sort(not_sorted_list[diviner:])
    res = merge(left, right)
    return res


def merge(list1, list2) -> list:
    res = []
    # print(list1, list2)
    i = 0
    j = 0
    for _ in range(len(list1)):
        if list1[i] <= list2[j]:
            res.append(list1[i])
            i += 1
        else:
            res.append(list2[j])
            j += 1
    # print(f"i: {i}, j: {j}")
    # print(f"res: {res}")
    if i <= len(list1) - 1:
        res = res + list1[i:]
    if j <= len(list2) - 1:
        res += list2[j:]
    # print(res)
    return res

    
not_sorted_list = [0, 0, -16, 2, 1, 3, 4, 6, 5, 7, 8,8]
res = merge_sort(not_sorted_list)
print(f"result is: {res}")