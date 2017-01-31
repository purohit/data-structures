# Definition for singly-linked list with a random pointer.
import unittest

class RandomListNode(object):
     def __init__(self, x):
         self.label = x
         self.next = None
         self.random = None

class Solution(object):
    def copyRandomList(self, head):
        oldToNew = {}
        prev, cur, newhead = None, None, None
        while head is not None:
            cur = RandomListNode(head.label)
            cur.random = head.random
            oldToNew[head] = cur
            if prev:
                prev.next = cur
            else:
                newhead = cur
            prev, head = cur, head.next
        cur = newhead
        while cur is not None:
            if cur.random:
                cur.random = oldToNew[cur.random]
            cur = cur.next
        return newhead

class TestSolution(unittest.TestCase):
    def test_empty(self):
        s = Solution()
        head = s.copyRandomList(None)
        self.assertIsNone(head)

    def test_blank(self):
        a = RandomListNode("A")
        s = Solution()
        head = s.copyRandomList(a)
        self.assertEquals(head.label, "A")
        self.assertIsNone(head.next)

    def test_simple(self):
        a = RandomListNode("A")
        b = RandomListNode("B")
        c = RandomListNode("C")
        d = RandomListNode("D")
        a.next = b
        b.next = c
        c.next = d
        a.random = c
        c.random = c
        d.random = a
        s = Solution()
        head = s.copyRandomList(a)
        self.assertEquals(head.label, "A")
        self.assertEquals(head.random.label, "C")
        self.assertEquals(head.next.label, "B")
        self.assertIsNone(head.next.random)
        self.assertEquals(head.next.next.label, "C")
        self.assertEquals(head.next.next.random.label, "C")
        self.assertEquals(head.next.next.next.label, "D")
        self.assertEquals(head.next.next.next.random.label, "A")
        self.assertIsNone(head.next.next.next.next)

if __name__ == '__main__':
    unittest.main()
