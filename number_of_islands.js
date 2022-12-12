// User function Template for javascript


class Pair {
    constructor(f,s) {
        this.first = f
        this.second = s
    }
}

/**
 * @param {string[][]} grid
 * @returns {number}
*/
class Solution {
    // Function to find the number of islands.
    numIslands(grid) {
        // code here
        let m = grid.length
        let n = grid[0].length
        let vis = new Array(m).fill(false).map(() => new Array(n).fill(false))
        
        
        
        let bfs = function(row, col, vis, grid) {
            let m = grid.length
            let n = grid[0].length
            let q = []
            vis[row][col] = true
            q.push(new Pair(row,col))
            
            while(q.length != 0) {
                let node = q.shift()
                let row = node.first
                let col = node.second
                
                for(let drow = -1; drow<=1; drow++) {
                    for(let dcol = -1; dcol<=1; dcol++) {
                        let nrow = drow + row
                        let ncol = dcol + col
                        if(ncol >= 0 && nrow >= 0 && ncol < n && nrow < m && grid[nrow][ncol] == '1' && vis[nrow][ncol] == false) {
                            vis[nrow][ncol] = true
                            q.push(new Pair(nrow, ncol))
                        }
                    }
                }
                
            }
        }
        
        let count = 0
        //console.log(grid)
        for(let row = 0; row < m; row++) {
            for(let col = 0; col < n; col++) {
                if(!vis[row][col] && grid[row][col] == '1') {
                    count++
                    bfs(row,col, vis, grid)
                }
            }
            //console.log(vis)
        }
        
        return count
        
        
    }
}
