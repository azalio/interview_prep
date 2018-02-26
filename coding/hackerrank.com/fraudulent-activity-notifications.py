#!/usr/bin/env python3
import statistics


# @profile
def activityNotifications(expenditure, window):
    # Complete this function
    notify, i = 0, 0
    expend_len = len(expenditure)
    while i + window < expend_len:
        # print(expenditure[i:window+i])
        mid = statistics.median(expenditure[i:window+i])
        print("mid: {}".format(mid))
        if expenditure[i+window] >= 2*mid:
            notify += 1
#            print("notify: {}".format(notify))
        i += 1
    return notify


if __name__ == "__main__":
    with open("fraudulent-activity-notifications.txt") as fh:
        n, d = fh.readline().split(" ")
        n, d = int(n), int(d)
        expenditure = fh.readline().split(" ")
        expenditure = list(map(int, expenditure))
    days = n
    window = d
    expenditure = expenditure[0:11000]

    days = 9
    window = 4
    expenditure = [1, 1, 4, 2, 3, 6, 8, 4, 5]


    notify = activityNotifications(expenditure, window)
    print(notify)

# kernprof -v -l ./fraudulent-activity-notifications.py
# 13
# Wrote profile results to fraudulent-activity-notifications.py.lprof
# Timer unit: 1e-06 s
#
# Total time: 2.68198 s
# File: ./fraudulent-activity-notifications.py
# Function: activityNotifications at line 5
#
# Line #      Hits         Time  Per Hit   % Time  Line Contents
# ==============================================================
#      5                                           @profile
#      6                                           def activityNotifications(expenditure, window):
#      7                                               # Complete this function
#      8         1        222.0    222.0      0.0      notify, i = 0, 0
#      9         1          5.0      5.0      0.0      expend_len = len(expenditure)
#     10      1002        749.0      0.7      0.0      while i + window < expend_len:
#     11                                                   # print(i+window)
#     12      1001    2676945.0   2674.3     99.8          mid = statistics.median(expenditure[i:window+i])
#     13      1001       3084.0      3.1      0.1          if expenditure[i+window] >= 2*mid:
#     14        13          7.0      0.5      0.0              notify += 1
#     15                                           #            print("notify: {}".format(notify))
#     16      1001        965.0      1.0      0.0          i += 1
#     17         1          1.0      1.0      0.0      return notify
