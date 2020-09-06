#!/usr/bin/env python3

def odd_even(l) -> list:
    even_pointer = 0 # O(1)
    odd_pointer = len(l) - 1 # O(1)
    for i in range(len(l)): # O(N)
        if even_pointer == odd_pointer: # O(1)
            break # O(1)
        if not l[even_pointer]%2: # O(1)
            # value is odd
            l[even_pointer], l[odd_pointer] = l[odd_pointer], l[even_pointer] # O(1)
            odd_pointer -= 1 # O(1)
        else:
            even_pointer += 1 # O(1)
    return l # O(1)

l = [1,2,3,4,5,5,6,7,7,8,9,10] # O(1)
res = odd_even(l) # O(1)

print(f"result is: {res}") # O(1)