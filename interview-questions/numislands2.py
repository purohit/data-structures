class Solution(object):
    class Node(object):
        def __init__(self):
            self.rank = 0
            self.prev = None

    def inBounds(self, r, c):
        return 0 <= r < self.m and 0 <= c < self.n

    def find(self, node):
        # Returns the representative. Optimization: path compression.
        if not node.prev: return node
        head = node
        while head.prev:
            head = head.prev
        # Path compression
        cur = node
        while node != head:
            tmp = node.prev
            node.prev = head
            node = tmp
        return head

    def merge(self, a, b):
        # Given two nodes we're guaranteed are in the union-set, merge them. Optimization: ranking.
        a, b = self.find(a), self.find(b)
        if a == b: return b
        self.islands -= 1 # When merging continuous land, # of islands must decrease.
        if a.rank <= b.rank:
            if a.rank == b.rank:
                b.rank += 1
            a.prev = b
            return b
        b.prev = a
        return a

    def addLand(self, r, c):
        """ Add land and return the new number of islands """
        # Out-of-bounds or already marked as land
        if not self.inBounds(r, c) or (r, c) in self.uf: return self.islands
        # Create a new island and mark the grid as land
        self.islands += 1
        cur = self.Node()
        self.uf[(r,c)] = cur
        # Merge it with any of in-bounds, land neighbors
        for nr, nc in [(r+1, c), (r-1, c), (r, c+1), (r, c-1)]:
            if not self.inBounds(nr, nc) or (nr, nc) not in self.uf:
                continue
            cur = self.merge(cur, self.uf[(nr, nc)])
        return self.islands

    def numIslands2(self, m, n, positions):
        """ Time complexity is O(k*alpha(mn)) ~= O(k) since
        alpha(x) <= 4 (constant) for any reasonable x, as it is the
        inverse Ackermann function """
        self.m, self.n = m, n
        self.uf = {} # Union-find data set
        self.islands = 0
        return [self.addLand(r, c) for r, c in positions]

if __name__ == '__main__':
    print Solution().numIslands2(3, 3, [[0,0], [0,1], [1,2], [2,1], [1, 1]])
    print Solution().numIslands2(0, 0, [[-1, 1]])
    print Solution().numIslands2(5, 5, [[0, 0], [0, 0]])
