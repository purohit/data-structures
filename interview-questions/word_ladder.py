import unittest

class Solution(object):
     def ladderLength(self, beginWord, endWord, wordList):
        return 0

class TestSolution(unittest.TestCase):
    def test_basic(self):
        s = Solution()
        words = set(["dog", "fog", "log", "pig", "fig", "ate", "big", "hog", "pin"])
        self.assertEquals(1, s.ladderLength("dog", "hog", words))
        self.assertEquals(4, s.ladderLength("dog", "pig", words))
        self.assertEquals(5, s.ladderLength("dog", "pin", words))
        self.assertEquals(5, s.ladderLength("pin", "dog", words))
        self.assertEquals(0, s.ladderLength("dog", "ate", words))


if __name__ == '__main__':
    unittest.main()
