import unittest


class Solution(object):
    def calc(self, s, LLi, LRi, Li, Ri):
        while (Li >= 0 and Ri < len(s)):
            if s[Li] != s[Ri]:
                break
            Li -= 1
            Ri += 1

        if Ri - Li - 2 > LRi - LLi:
            LLi = Li + 1
            LRi = Ri - 1

        return (LLi, LRi)

    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        sLen = len(s)
        if sLen == 1:
            return s

        LLi = 0
        LRi = 0

        for key, _ in enumerate(s):
            if sLen - (LRi - LLi) / 2 < key:
                break

            LLi, LRi = self.calc(s, LLi, LRi, key, key)

            if key < 1:
                continue

            LLi, LRi = self.calc(s, LLi, LRi, key - 1, key)

        return s[LLi:LRi + 1]


class SolutionTest(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_solution(self):
        s = "abababab"
        expected = "abababa"
        self.assertEqual(expected, self.solution.longestPalindrome(s))


if __name__ == '__main__':
    unittest.main()
