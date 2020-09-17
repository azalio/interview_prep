#!/usr/bin/env python3


class Solution:
    def timeDifference(self, times): # O(n**2)
        '''
        :type times: list of str
        :rtype: int
        '''
        # times = ["00:03", "23:59", "12:03"]
        times = list(map(lambda x: x.split(':'), times)) # O(n)
        for i, v in enumerate(times): # O(n)
            times[i] = int(v[0])*60+int(v[1])
            # print(f"times[i]: {times[i]}")
            if times[i] <= 720: # O(1)
                times[i] = times[i] + 1440 # O(1)
        # print(f"times: {times}")
        index = 0 # O(1)
        min_seen = 1441 # O(1)
        while index < len(times) - 1: # O(n**2) = n(n+1)/2 = (n-1 + n-2 + n-3 ...)
            # print(f"min_seen: {min_seen}")
            for i in times[index + 1:]: # O(n)
                if min_seen > abs(times[index] - i): # O(1)
                    min_seen = abs(times[index] - i) # O(1)
            index += 1 # O(1)
        return min_seen # O(1)


    def timeDifference2(self, times):
        '''
        :type times: list of str
        :rtype: int
        '''
        seen = [False] * (24 * 60) # Create a list on False for every min

        for time in times:
            n = self.time_to_int(time) # Transform to minutes

            if seen[n]:
                return 0

            seen[n] = True

        prev = -1
        min_diff = (24 * 60) + 1
        first = -1
        for i in range(0, len(seen)):
            if seen[i]:
                if prev != -1:
                    min_diff = min(min_diff, i - prev)
                else:
                    first = i

                prev = i
                print(f"first: {first}, prev: {prev}, min_diff: {min_diff}")

        # Wrap-around check
        min_diff = min(min_diff, first + 1440 - prev)

        return min_diff

    def time_to_int(self, time):
        # time format - HH:mm
        hours = time[0:2]
        minutes = time[3:5]

        return int(hours) * 60 + int(minutes)

times = ["00:03", "23:59", "12:03"]
obj = Solution()
res = obj.timeDifference(times)
assert res == 4

times = ["00:01","00:02"]
obj = Solution()
res = obj.timeDifference(times)
assert res == 1


times = ["23:59", "00:03", "12:03"]
obj = Solution()
res = obj.timeDifference2(times)
assert res == 4

times = ["00:01","00:02"]
obj = Solution()
res = obj.timeDifference2(times)
assert res == 1

# print(f"min diffirence is: {res}")