from collections import deque

class Solution(object):
    def shortestDistance(self, grid):
	if not grid or not grid[0]:
	    return -1
	rows, cols = len(grid), len(grid[0])
        building_spots = [(r, c) for r in range(rows) for c in range(cols) if grid[r][c] == 1]
        total_buildings = len(building_spots)
        hits, distTotal = [[0]*cols for _ in range(rows)], [[0]*cols for _ in range(rows)]

	def search(r, c):
	    seen = [[False]*cols for _ in range(rows)]
	    seen[r][c] = True
            buildings_hit = 1
            queue = deque([(r, c, 0)])
	    while queue:
		r, c, dist = queue.pop()
                for nr, nc in ((r+1, c), (r, c+1), (r-1, c), (r, c-1)):
                    if nr < 0 or nc < 0 or nr >= rows or nc >= cols or seen[nr][nc]:
                        continue
                    seen[nr][nc] = True
                    if grid[nr][nc] == 0: # Found a traversable spot (house)
                        queue.appendleft((nr, nc, dist+1))
                        hits[nr][nc] += 1 # We can hit this building; important later!
                        distTotal[nr][nc] += dist + 1
                    elif grid[nr][nc] == 1:
                        buildings_hit += 1
	    return buildings_hit == total_buildings

        for r, c in building_spots:
            if not search(r, c):
                return -1

        mindist = -1
        for r in range(rows):
            for c in range(cols):
                # Not a plot of land, or isn't reachable by all buildings
                if grid[r][c] != 0 or hits[r][c] != total_buildings:
                    continue
                if mindist == -1 or distTotal[r][c] < mindist:
                    mindist = distTotal[r][c]
        return mindist

if __name__ == '__main__':
    s = Solution()
    grid = [
            [1,0,2,0,1],
            [0,0,0,0,0],
            [0,0,1,0,0]
            ]
    print s.shortestDistance(grid)
    grid = []
    print s.shortestDistance(grid)
    grid = [[1],[0]]
    print s.shortestDistance(grid)
    grid = [[0,1,2,0,2,0,2,1,2,0,0,2,2,2,2,0,2,0,2,0,2,2,2,2,0],
            [2,0,0,0,1,0,0,0,2,2,0,0,0,2,2,0,0,2,2,0,2,0,0,0,2],
            [0,2,0,2,2,0,0,0,0,2,2,0,2,0,0,2,0,0,0,2,0,2,0,0,2],
            [0,0,0,0,0,0,0,0,0,2,2,2,2,2,0,0,2,1,1,0,0,0,2,0,2],
            [2,2,0,0,2,0,0,0,0,0,0,1,2,2,0,2,2,2,0,0,0,1,2,2,0],
            [0,2,0,0,2,0,0,0,2,0,0,2,0,2,0,0,1,1,0,0,0,0,2,0,2],
            [0,0,0,2,2,2,2,0,2,0,0,2,0,2,2,0,0,0,0,0,0,2,0,2,0],
            [0,0,2,0,2,1,2,0,0,0,2,2,0,0,0,0,2,0,0,2,2,0,2,0,0],
            [0,0,0,0,0,1,1,2,0,2,0,0,2,2,2,0,0,2,2,0,0,0,0,2,0],
            [0,0,0,2,0,2,2,0,2,2,2,0,0,0,0,2,0,0,2,0,0,0,2,2,0],
            [0,2,2,2,2,0,0,2,2,2,0,0,2,2,0,0,0,2,0,1,2,0,0,1,0],
            [0,0,2,0,0,2,0,0,2,0,0,0,0,0,2,2,0,0,0,1,0,0,0,2,0],
            [0,2,2,0,0,2,2,0,0,2,2,2,2,0,0,0,2,0,2,2,0,2,0,0,0],
            [0,0,0,0,0,0,0,2,0,0,2,0,0,0,2,0,0,0,0,0,0,0,0,0,0],
            [0,0,0,0,0,2,0,0,0,0,0,2,0,2,0,0,0,0,1,0,0,0,2,2,0],
            [2,0,2,0,0,0,2,0,2,2,0,0,0,0,0,0,0,0,2,0,0,2,2,2,1],
            [0,2,0,0,0,2,2,2,0,2,0,0,2,2,2,0,0,0,2,0,2,0,2,2,0],
            [0,0,0,2,1,2,0,0,0,0,2,0,0,0,2,0,2,0,0,0,2,2,2,0,2],
            [0,2,0,0,0,0,2,2,2,0,0,0,2,2,0,0,0,2,0,0,0,0,0,2,0],
            [1,0,0,0,0,0,0,0,0,0,0,2,0,2,0,0,0,2,2,0,2,2,0,2,0],
            [0,2,1,0,1,0,2,2,0,0,0,0,0,0,2,0,2,2,2,2,0,2,2,0,0],
            [2,0,0,0,0,0,2,2,0,2,0,2,2,0,0,0,0,0,2,2,2,2,0,0,0],
            [0,2,2,0,2,0,2,2,0,2,0,0,2,2,0,2,0,2,0,2,0,0,1,0,2],
            [0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2,0,2,2,2,2,0,0,0],
            [0,0,0,0,0,2,1,0,2,0,2,0,0,0,0,2,2,0,2,2,0,0,2,0,2],
            [2,2,0,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,2,2,0,0,0],
            [0,0,0,2,0,0,2,0,0,0,0,2,2,0,0,2,1,2,0,2,0,2,2,1,2],
            [2,2,0,2,0,0,0,0,0,0,0,2,2,0,0,0,0,0,0,0,2,0,0,0,2]]
    print s.shortestDistance(grid)

