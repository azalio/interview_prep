#!/usr/bin/python3

def sum_of_two(l: list, s: int) -> list: # O(n)
    compiment = dict() # O(1)
    result = [] # O(1)
    for i in l: # O(n)
        comp = s - i # O(1)
        if compiment.pop(comp, None) != None: # O(1)
            result.append([i, comp]) # O(1)
        else:
            compiment[i] = True # O(1)

    return result # O(1)


l = [-1,1,0,32,2,5, -4, -5, 2,-32,-2]
s = 1
sum_of_two(l, s)
assert sum_of_two(l, s) == [[0, 1], [2, -1], [-4, 5]]

l = [2,7,11,15]
s = 9
assert sum_of_two(l, s) == [[7, 2]]