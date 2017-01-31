/**
*
*  * @param {number[][]} grid
*
*   * @return {number}
*
*    */
var shortestDistance = function(grid) {
    var empty = 0;
    var building = 1;
    var obstacle = 2;

    var ma
    // Let's do BFS for any open spot that we find
    for (var r = 0; r < grid.length; r += 1) {
        for (var c = 0; c < grid.length[r]; c += 1) {
            if grid[r][c] != empty {
                continue
            }



        }
    }
};

var res = shortestDistance([[1,0,2,0,1],[0,0,0,0,0],[0,0,1,0,0]]);
console.log(res);
