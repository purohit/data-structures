import unittest

# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def longestConsecutive(self, root):
        if root == None:
            return 0
        best = 0
        stack = [(root, 1)]
        while stack:
            cur, run = stack.pop()
            best = max(run, best)
            for x in [n for n in [cur.left, cur.right] if n]:
                val = 1 if x.val - cur.val == 1 else 0
                stack.append((x, run+val))
        return best


class TestTree(unittest.TestCase):
    def test_one(self):
        x = TreeNode(2)
        x.right = TreeNode(3)
        x.right.left = TreeNode(2)
        x.right.left.left = TreeNode(1)
        s = Solution()
        self.assertEquals(s.longestConsecutive(x), 2)

    def test_sub(self):
        head = TreeNode(2)
        head.left = TreeNode(1)
        head.left.left = TreeNode(0)
        head.left.right = TreeNode(2)
        head.left.right.right = TreeNode(3)
        head.left.right.right.right = TreeNode(4)
        head.right = TreeNode(4)
        head.right.right = TreeNode(5)
        head.right.right.right = TreeNode(6)
        head.right.right.right.left = TreeNode(0)
        s = Solution()
        self.assertEquals(s.longestConsecutive(head), 4)

    def test_simple(self):
        x = TreeNode(2)
        s = Solution()
        self.assertEquals(s.longestConsecutive(x), 1)

    def test_simp_easy(self):
        x = TreeNode(1)
        x.right = TreeNode(2)
        s = Solution()
        self.assertEquals(s.longestConsecutive(x), 2)

    def test_none(self):
        s = Solution()
        self.assertEquals(s.longestConsecutive(None), 0)

if __name__ == '__main__':
    unittest.main()

