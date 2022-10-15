//User function Template for javascript

function solve(board) {
    for(let i=0; i<board.length; i++) {
        for(let j=0; j<board[0].length; j++) {
            if(board[i][j] == 0) {
                for(let c = 1; c<=9; c++) {
                    if(isValid(board, i,j, c)) {
                        board[i][j] = c
                        if(solve(board)) {
                            return true
                        }
                        board[i][j] = 0
                    }
                }
                return false
            } 
        }
        
    }
    return true
}

function isValid(board, row,col, c) {
    for(let i =0; i<9; i++) {
        //check row
        if(board[row][i] == c) return false
        //check column
        if(board[i][col] == c) return false

        //check sub square
        if(board[3 * ~~(row/3) + ~~(i/3)][3* ~~(col/3) + i % 3] == c) return false
    }
    return true
}


class Solution {    
    /**
     * @param {number[][]} grid
     * @returns {boolean}
     */
    
    SolveSudoku(grid)
    {
        //your code here
        if(solve(grid)) {
            this.printGrid(grid)
            return true
        } else {
            return false
        }
        
        
    }
    
    
    /**
     * @param {number[][]} grid
     */
    
    printGrid(grid)
    {
        //your code here
        for(let i=0; i<grid.length; i++) {
            for(let j=0; j<grid[0].length; j++) {
                process.stdout.write(grid[i][j] + " ")
            }
            //console.log()
        }
        
    }

    
}



let s = new Solution()
let grid = 
    [
     [3, 0, 6, 5, 0, 8, 4, 0, 0],
     [5, 2, 0, 0, 0, 0, 0, 0, 0],
     [0, 8, 7, 0, 0, 0, 0, 3, 1],
     [0, 0, 3, 0, 1, 0, 0, 8, 0],
     [9, 0, 0, 8, 6, 3, 0, 0, 5],
     [0, 5, 0, 0, 9, 0, 6, 0, 0],
     [1, 3, 0, 0, 0, 0, 2, 5, 0],
     [0, 0, 0, 0, 0, 0, 0, 7, 4],
     [0, 0, 5, 2, 0, 6, 3, 0, 0]
    ]

s.SolveSudoku(grid)
