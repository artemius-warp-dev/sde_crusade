//User function Template for javascript



/**
 * @param {string} str
 * @returns {number}
*/

class Solution{
    countMin(str){
        //code here
          //code here
        let s1 = str
        let s2 = s1.split("").reverse().join("")
        let n = s1.length
        let dp = new Array(n+1).fill().map(() => new Array(n+1).fill(-1))
        let f = function(i,j, s1,s2, dp) {
            if(i == 0 || j == 0) return 0
            if(dp[i][j] != -1) return dp[i][j]
            if(s1[i-1] == s2[j-1]) {
                dp[i][j] = 1 + f(i-1,j-1,s1,s2, dp)
            } else {
                dp[i][j] = Math.max(f(i-1, j, s1, s2, dp), f(i, j-1, s1,s2,dp))
            }
            return dp[i][j]
        }
        return str.length - f(n,n,s1,s2,dp)
        
    }
}
