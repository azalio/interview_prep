#!/usr/bin/env python3

class Solution:
    def timeDifference(self, times):
        '''
        :type times: list of str
        :rtype: int
        '''
        # times = ["00:03", "23:59", "12:03"]
        times = list(map(lambda x: x.split(':'), times))
        for i, v in enumerate(times):
            times[i] = int(v[0])*60+int(v[1])
            if times[i] > 720:
                times[i] = 1440 - times[i]
        print(f"times: {times}")
        times.sort()
        return times[0] + times[1]

times = ["00:03", "23:59", "12:03"]
obj = Solution()
res = obj.timeDifference(times)
assert res == 4

times = ["00:01","00:02"]
obj = Solution()
res = obj.timeDifference(times)
assert res == 1

# print(f"min diffirence is: {res}")