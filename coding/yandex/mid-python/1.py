class Solution:
    @staticmethod
    def sumTwo(nums, target):
        seen = {}
        result = []
        for i, v in enumerate(nums):
            tmp = []
            remaning = target - v
            if remaning in seen:
                tmp.append(remaning)
                tmp.append(v)
                result.append(tmp)
            seen[v] = i
        return result

    def threeSum(self, nums: List[int]) -> List[List[int]]:
        result = []
        for i, v in enumerate(nums):
            tsum_result = self.sumTwo(nums[:i] + nums[i+1:], 0 - v)
            # print(tsum_result)
            if tsum_result:
                [x.append(v) for x in tsum_result]
                tsum_result = [sorted(x) for x in tsum_result]
                # print(tsum_result)
                # tsum_result.sort()
                for i in tsum_result:
                    if i in result:
                        continue
                    result.append(i)
        result = sorted(result)
        return result