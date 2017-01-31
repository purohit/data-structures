from collections import OrderedDict
"""
Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and set.
get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
set(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.
"""

class LRUCache(object):
    def __init__(self, capacity):
        self.capacity = capacity
        self.items = OrderedDict()
        self.timestep = 0

    def get(self, key):
        if key not in self.items:
            return -1
        val = self.items.pop(key)
        self.items[key] = val
        return val

    def set(self, key, value):
        if key in self.items:
            self.items.pop(key)
            self.items[key] = value
            return
        # If at capacity, remove the least frequent value
        if len(self.items) == self.capacity:
            self.items.popitem(last=False)
        # Add the latest value
        self.items[key] = value

import unittest

class TestLRUCache(unittest.TestCase):
    def test_simple(self):
        l = LRUCache(2)
        self.assertEquals(l.get(5), -1)
        l.set(5, 10)
        self.assertEquals(l.get(5), 10)
        l.set(1, 7)
        self.assertEquals(l.get(1), 7)
        l.set(2, 3)
        self.assertEquals(l.get(2), 3)
        self.assertEquals(l.get(5), -1)
        self.assertEquals(l.get(1), 7)
        l.set(0, 5)
        self.assertEquals(l.get(2), -1)

    def test_leetcode(self):
        l = LRUCache(2)
        l.set(2, 1)
        l.set(2, 2)
        self.assertEquals(l.get(2), 2)
        l.set(1, 1)
        l.set(4, 1)
        self.assertEquals(l.get(2), -1)

    def test_leetcode_2(self):
        l = LRUCache(2)
        l.set(2, 1)
        l.set(1, 2)
        l.set(2, 3)
        l.set(4, 1)
        self.assertEquals(l.get(1), -1)
        self.assertEquals(l.get(2), 3)

if __name__ == '__main__':
    unittest.main()
