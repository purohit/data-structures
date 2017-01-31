import unittest

class NumMatrix(object):
    def __init__(self, matrix):
        self.valid = True
        self.matrix = matrix
        self.rows = len(self.matrix)
        if self.rows == 0 or len(self.matrix[0]) == 0:
            self.valid = False
            return
        self.cols = len(self.matrix[0])
        for i in range(self.rows):
            acc = 0
            for j in range(self.cols):
                acc += matrix[i][j]
                matrix[i][j] = acc

    def update(self, row, col, val):
        # Out of bounds
        if not self.valid or row < 0 or row >= self.rows or col < 0 or col >= self.cols:
            return
        prev = 0
        if col > 0:
            prev = self.matrix[row][col-1]
        delta = val - self.matrix[row][col] + prev
        for c in range(col, self.cols):
            self.matrix[row][c] += delta

    def sumRegion(self, row1, col1, row2, col2):
        # Out of bounds
        if not self.valid or row1 < 0 or row2 >= self.rows or col1 < 0 or col2 >= self.cols:
            return 0
        acc = 0
        for r in range(row1, row2+1):
            if col1 < 1:
                acc += self.matrix[r][col2]
            else:
                acc += self.matrix[r][col2] - self.matrix[r][col1-1]
        return acc


class TestMatrix(unittest.TestCase):
    def test_bad(self):
        matrix = []
        n = NumMatrix(matrix)
        self.assertEquals(n.sumRegion(2, 1, 4, 3), 0)
        n.update(3, 2, 2)
        self.assertEquals(n.sumRegion(2, 1, 4, 3), 0)

    def test_good(self):
        matrix = [
                [3, 0, 1, 4, 2],
                [5, 6, 3, 2, 1],
                [1, 2, 0, 1, 5],
                [4, 1, 0, 1, 7],
                [1, 0, 3, 0, 5]
                ]
        n = NumMatrix(matrix)
        self.assertEquals(n.sumRegion(2, 1, 4, 3), 8)
        n.update(3, 2, 2)
        self.assertEquals(n.sumRegion(2, 1, 4, 3), 10)

    def test_edge(self):
        matrix = [
                [3, 0, 1, 4, 2],
                [5, 6, 3, 2, 1],
                [1, 2, 0, 1, 5],
                [4, 1, 0, 1, -7],
                [1, 0, 3, 0, 5]
                ]
        n = NumMatrix(matrix)
        n.update(4, 4, 17)
        self.assertEquals(n.sumRegion(3, 2, 4, 4), 14)

if __name__ == '__main__':
    unittest.main()
