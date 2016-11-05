import unittest


class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        pool = {}
        for index, n in enumerate(nums):
            if n in pool:
                return [pool[n], index]
            pool[target - n] = index


class SolutionTest(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_solution(self):
        nums = [2, 5, 7, 11, 15]
        target = 9
        result = self.solution.twoSum(nums, target)
        expected = [0, 2]
        self.assertEqual(expected, result)


if __name__ == '__main__':
    unittest.main()
