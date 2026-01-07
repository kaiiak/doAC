# https://leetcode.cn/problems/happy-number/
class Solution:
    def isHappy(self, n: int) -> bool:
        return self.happySum(n, set({}))[1]

    def happySum(self, n: int, cache: set) -> tuple[int, bool]:
        if n in cache:
            return n, False
        else:
            cache.add(n)
        sum = 0
        for x in str(n):
            i = int(x)
            sum += (i ** 2)
        if sum != 1:
            return self.happySum(sum, cache)
        else:
            return sum, True
