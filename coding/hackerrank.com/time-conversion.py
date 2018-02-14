#!/usr/bin/env python3

import sys

def timeConversion(s):
    # Complete this function
    arr = s.split(':')
    if arr[-1].endswith('AM'):
        if arr[0] == '12':
            arr[0] = '00'
        arr[-1] = arr[-1].rstrip('AM')
    else:
        arr[-1] = arr[-1].rstrip('PM')
        if arr[0] == '12':
            pass
        else:
            arr[0] = str(int(arr[0]) + 12)
        
    return ":".join(arr)

s = '12:00:00PM'
result = timeConversion(s)
print(result)
