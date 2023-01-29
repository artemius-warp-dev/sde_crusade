 class Node {
    constructor(r,c) {
        this.row = r
        this.col = c
    }
}


//User function Template for javascript

/**
 * @param {number[][]} grid
 * @returns {boolean}
*/
class Solution
{
    //Function to find whether a path exists from the source to destination.
    is_Possible(grid)
    {
        // code here
        let m = grid.length
        let n = grid[0].length
        
        
        let bfs = function(grid, m, n, s, d) {
            let q = []
            q.push(new Node(s.row, s.col))
        
            let dr = [-1, 0, 1, 0]
            let dc = [0, 1, 0, -1]
                
            while(q.length != 0) {
                //console.log(pq.harr[0])
                let node = q.shift()
                
                let r = node.row
                let c = node.col
                
                for(let i =0; i<4; i++) {
                    let nrow = r + dr[i]
                    let ncol = c + dc[i]
                    
                    
                    if(nrow >=0 && ncol >= 0 && nrow < m && ncol < n && grid[nrow][ncol] != 0 ) {
                        if(grid[nrow][ncol] == 2)
                            return true
                        q.push(new Node(nrow, ncol))
                        grid[nrow][ncol] = 0
                    } else {
                        continue
                    }
                }
                
               
            }
            return false
        }
        
        
        
        
        let s = -1, d = -1
        for(let i=0; i<m; i++) {
            for(let j=0; j<n; j++) {
               if(grid[i][j] == 1) 
                    s = new Node(i,j)
               if(grid[i][j] == 2)
                    d = new Node(i,j)
            }
        }
        
        //console.log(s,d)
        //return 0
        return bfs(grid, m,n, s, d)
 
        
        
    }
}
