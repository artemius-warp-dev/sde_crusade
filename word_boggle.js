//User function Template for javascript
1
/**
 * @param {string[][]} board
 * @param {string[]} dictionary
 * @return {string[]}
 */

class Solution {


 
    
    dfs(board, visited, word, row, col, i, j, idx) {


        
        
        if(idx > word.length) return false

        if(idx == word.length) return true
        
        if(i >= row || j >= col || i < 0 || j < 0) return false

       

        if (word[idx] != board[i][j]) return false

        

         if(visited[i][j] == 1) return false

        
        visited[i][j] = 1
        let up = this.dfs(board, visited, word, row, col, i - 1, j, idx+1) 
        
        let down = this.dfs(board, visited, word, row, col, i + 1, j, idx+1)
        let right = this.dfs(board, visited, word, row, col, i, j + 1, idx+1)
        let left = this.dfs(board, visited, word, row, col, i, j-1, idx+1)
        let up_right = this.dfs(board, visited, word, row, col, i-1, j+1, idx+1)
        let up_left = this.dfs(board, visited, word, row, col, i-1, j-1, idx+1)
        let down_right = this.dfs(board, visited, word,  row, col, i+1, j+1, idx+1)
        let down_left = this.dfs(board, visited, word, row, col, i+1, j-1, idx+1)
        visited[i][j] = '*'


        return up || down || right || left || up_right || up_left || down_right || down_left
        
    }
    


    wordBoggle(board,dictionary){
        //code here
        let row = board.length, col = board[0].length
        let visited = new Array(row)
        let result = new Set()
        
        
        for(let i = 0; i < row; i++) {
            visited[i] = new Array(col).fill('*')
        }

        for(let s = 0; s<dictionary.length; s++) {
            let word = dictionary[s]
            for(let i = 0; i < row; i++) {
                for(let j=0; j<col; j++) {
                    if(this.dfs(board, visited, word, row, col, i,j, 0)) {
                        result.add(word)
                    }
                }

            }
        }

        return Array.from(result).sort()
    }
}


let s = new Solution()
let board = [['d','d'],
             ['b','f'],
             ['e','c'],
             ['b', 'c'],
             ['d', 'c']]
let dictionary = ["bcd", "db"]
console.log(s.wordBoggle(board, dictionary))
