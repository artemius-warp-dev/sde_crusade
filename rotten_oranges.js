//User function Template for javascript


class Node {
    constructor() {
        this.pair = []
        this.time = null
        this.prev = null
        this.next = null
    }
}

class Deque {
    constructor() {
        this.head = null
        this.tail = null
    }

    append(pair, time) {
        let temp = new Node()
        temp.pair = pair
        temp.time = time
        temp.prev = temp.next = null

        if(this.head == null) {
            this.head = this.tail = temp
            return
        }

        this.tail.next = temp
        temp.prev = this.tail
        this.tail = temp
    }


    remove(node) {
        if(this.head == node)
            this.head = this.head.next
        if(this.tail == node)
            this.tail = this.tail.prev
        if(node.next != null)
            node.next.prev = node.prev
        if(node.prev != null)
            node.prev.next = node.next
    }
    
}



/**
 * @param {number[][]} grid
 * @returns {number}
*/
class Solution
{
    //Function to find minimum time required to rot all oranges. 
    orangesRotting(grid) 
    {
        // code here
        let d = new Deque()
        let m = grid.length
        let n = grid[0].length
        let visited = new Array(m)
        for(let i=0; i<m; i++) {
            visited[i] = new Array(n)
            for(let j=0; j<n; j++) {
                visited[i][j] = 0
                if(grid[i][j] == 2) {
                    d.append([i,j], 0)
                    visited[i][j] = 2
                } else
                    visited[i][j] = 0
                
            }
        }

        let tm = 0
        let drow = [0, 1, 0, -1]
        let dcol = [-1, 0, 1, 0]
        while(d.head != null) {
            let t = d.head.time
            let r = d.head.pair[0]
            let c = d.head.pair[1]
            tm = Math.max(tm, t)
            d.remove(d.head)
            for(let i = 0; i<4; i++) {
                let nrow = r + drow[i]
                let ncol = c + dcol[i]
                if(nrow >= 0 && ncol >= 0 && nrow < m
                   && ncol < n && visited[nrow][ncol] != 2
                   && grid[nrow][ncol] == 1) {
                    d.append([nrow, ncol], t + 1)
                    visited[nrow][ncol] = 2
                }
            }

        }

        for(let i = 0; i<m; i++)
            for(let j = 0; j<n; j++) {
                if(visited[i][j] != 2 && grid[i][j] == 1)
                    return -1
            }
        return tm
    }
}


let s = new Solution()
let grid = [[0,1,2],[0,1,2],[2,1,1]]
console.log(s.orangesRotting(grid))
