#!/usr/bin/python3
import time

def find_first_occ(l: list, n: int, start: int, end: int) -> int: # O(n) > O(log2(n)) > O(1) -> O(n)
    index = bin_search(l, start, end, n) # O(log2(n))
    if index == -1: # O(1)
        return -1 # O(1)
    for i in range(index + 1, -1, -1): # O(n)
        # print(f"l[{i}]: {l[i]}")
        if l[i] != n: # O(1)
            index = i + 1 # O(1)
            break # O(1)
        else:
            index = i # O(1)
    # print(f"index: {index}")
    return index # O(1)
    
def find_first_occ2(l: list, n: int, start: int, end: int) -> int:
    # O(log2(n)) повторенный O(log2(n)) раз.
    # O(log2(n)*log2(n)) -> O(2log2(n)) -> O(log2(n))
    index = bin_search(l, start, end, n) # O(log2(n))
    if index == -1 and (start == 0 and end == len(l)): # O(1)
        return -1 # O(1)
    # print(f"index: {index}")
    if index == 0: # O(1)
        return index # O(1)
    if l[index - 1] != l[index]: # O(1)
        return index # O(1)
    return find_first_occ2(l, n, 0, index) # O(1)


def bin_search(l: list, s: int, e: int, value: int) -> int: # O(log2(n))
    m = s + (e - s)//2 # O(1)
    # print(f"s: {s}, e: {e}, m: {m}")
    if l[m] == value: # O(1)
        return m # O(1)
    if m == s: # O(1)
        return s if l[s] == value else -1 # O(1)
    if l[m] > value: # O(1)
        e = m # O(1)
        return bin_search(l, s, e, value)
    else:
        s = m # O(1)
        return bin_search(l, s, e, value)


l = [1,2,3,4,4,4,4,4,5,6,7,8,9]
a = [0, 1, 2, 3, 4, 5]
assert bin_search(a, 0, len(a), -1) == -1
assert bin_search(a, 0, len(a), 0) == 0
assert bin_search(a, 0, len(a), 2) == 2
assert bin_search(a, 0, len(a), 5) == 5
assert bin_search(a, 0, len(a), 10) == -1
assert find_first_occ(l, 4, 0, len(l)) == 3
assert find_first_occ2(l, 4, 0, len(l)) == 3


A = [4]*10**7
# A = [1,2,3,4,4,4,5,6]
start = time.time()
res = find_first_occ(A, 4, 0, len(A))
end = time.time()
print(f"res: {res}")
print(f"func find_first_occ bin + linear time: {end - start}")

start = time.time()
res = find_first_occ2(A, 4, 0, len(A))
end = time.time()
print(f"res: {res}")
print(f"func find_first_occ2 bin + bin time: {end - start}")