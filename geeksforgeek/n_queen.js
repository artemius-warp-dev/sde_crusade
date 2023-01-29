//User function Template for javascript

/**
 * @param {number} n
 * @return {number[][]}
 */

class Solution {
    nQueen(n){
        //code here
        let board = new Array(n)
        let ans = []
        for(let i = 0; i<n; i++) {
            board[i] = new Array(n).fill('.')
        }

        let solve = function(col, board, ans, n) {
            if(col == n) {
                //console.log(board)
                let v = []
                for(let i =0; i<n; i++) {
                    for(let j = 0; j< n; j++) {
                        if(board[i][j] == 'Q') {
                            v.push(j+1)
                            //break
                        }
                    }
                   
                }
                ans.push(v)
                return 
            }

            for(let row =0; row < n; row++) {
                if(isSafe(board, row, col)) {
                    board[row][col] = 'Q'
                    solve(col+1, board, ans, n)
                    board[row][col] = '.'
                }
            }   

            
        }

        let isSafe = function(board, row, col) {
            let crow = row
            let ccol = col

            
            while(row >=0 && col >= 0) {
                if(board[row][col] == 'Q') return false
                row-=1
                col-=1
            }

            row = crow
            col = ccol

            while(col >= 0) {
                if(board[row][col] == 'Q') return false
                col-=1
            }

            row = crow
            col = ccol

            while(row < n && col >= 0) {
                if(board[row][col] == 'Q') return false
                row+=1
                col-=1
            }

            return true
            
            
        }

        solve(0, board, ans, n)
        //console.log(ans.sort())
        console.log(ans)
        return ans.sort((a, b) => a.reduce(function(c, v, i) {
            //console.log(c, v, b[i])
            c = c ? c : v - b[i]
            //console.log(c)
            return c
        }, 0
        
        ));
        

       
        
    }
}

let s = new Solution()
let n = 4
console.log(s.nQueen(n))
