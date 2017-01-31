import unittest

# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def pathTo(self, root, x):
        parents = {root: None}
        stack = [root]
        while len(stack) > 0:
            node = stack.pop()
            if node == x:
                break
            if node.left:
                parents[node.left] = node
                stack.append(node.left)
            if node.right:
                parents[node.right] = node
                stack.append(node.right)
        path = [x]
        cur = x
        while cur != None:
            path.append(parents[cur])
            cur = parents[cur]
        return path

    def lowestCommonAncestor(self, root, p, q):
        ppath = self.pathTo(root, p)
        qpath = self.pathTo(root, q)
        common = None
        while ppath and qpath:
            if ppath[-1] == qpath[-1]:
                common = ppath[-1]
                ppath = ppath[:-1]
                qpath = qpath[:-1]
                continue
            break
        return common

class TestLCA(unittest.TestCase):
    def test_simple(self):
        s = Solution()
        x = TreeNode(5)
        self.assertEquals(s.lowestCommonAncestor(x, x, x), x)
        x.left = TreeNode(6)
        x.right = TreeNode(7)
        self.assertEquals(s.lowestCommonAncestor(x, x.left, x.right), x)
        self.assertEquals(s.lowestCommonAncestor(x, x.right, x.left), x)
        x.left.left = TreeNode(8)
        x.left.right = TreeNode(9)
        x.right.right = TreeNode(10)
        self.assertEquals(s.lowestCommonAncestor(x, x.right, x.left.left), x)
        self.assertEquals(s.lowestCommonAncestor(x, x.left.right, x.left.left), x.left)

if __name__ == '__main__':
    unittest.main()
