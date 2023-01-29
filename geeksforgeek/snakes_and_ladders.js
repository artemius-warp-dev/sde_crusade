//User function Template for javascript

/**
 * @param {number[]} arr
 * @param {number} n
 * @returns {number}
*/

class Solution {
    
    minThrow(arr, n)
    {
        //arr.reverse()
        let vis = new Array(31).fill(false)
        let board = new Array(31).fill(-1)
        
        for(let i = 0; i<=n; i++) {
            board[arr[2*i]] = arr[2*i+1]
        }
        let q = []
        q.push([1,0])
        while (q.length != 0) {
                let qsize = q.length
                    let [square, moves] = q.shift()
                    if(vis[square]) continue
                    vis[square] = true
                    if(square == 30) return moves
                
                    for(let i=1; i<=6 && square + i <= 30; i++) {
                        let nextSquare = square + i
              
                        if(board[nextSquare] != -1) nextSquare = board[nextSquare]
                        if(!vis[nextSquare]) q.push([nextSquare, moves + 1])
     
                    }  
 
        }
        return -1
    }
}
