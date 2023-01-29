//User function Template for javascript

/**
 * @param {number[][]} Matrix
 * @param {number} n
 * @return {number}
 */
 

class Solution {
    maximumPath(Matrix,n){
        //code here
        let dp = new Array(n).fill().map(() => new Array(n).fill(-1))
       
        
        let f = function(r,c, matrix, dp) {
            if(c < 0 || c >= n || r == n) return Number.MIN_VALUE
            
            if(dp[r][c] != -1) return dp[r][c]
            let max = Number.MIN_VALUE
            for(let i=-1; i<=1; i++) {
                max = Math.max(max, matrix[r][c] + f(r+1, c+i, matrix, dp))
            }
            
            return dp[r][c] = max
        }
        
        let ans = Number.MIN_VALUE
        for(let i = 0; i<n; i++) {
            ans = Math.max(ans, f(0,i,Matrix, dp))
        }
        
        
        return ans
        
    }
}
