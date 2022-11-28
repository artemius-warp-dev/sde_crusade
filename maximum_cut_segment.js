//User function Template for javascript

/**
 * @param {number} n
 * @param {number} x
 * @param {number} y
 * @param {number} z
 * @returns {number}
*/

class Solution
{
    //Function to find the maximum number of cuts.
    maximizeTheCuts(n, x, y, z)
    {
        // code here
        let dp = new Array(n+1).fill(-1)
        let cuts = [x,y,z]
        let f = function(n, cuts, dp) {
            if(n==0) return 0
            if(n < 0) return Number.MIN_VALUE
            if(dp[n] != -1) return dp[n]
            let max = Number.MIN_VALUE
            for(let i=0; i<cuts.length; i++) {
                if(n-cuts[i] >= 0) {
                  let temp = f(n-cuts[i], cuts, dp)
                  if(temp != Number.MIN_VALUE) {
                     max = Math.max(max, temp+1) 
                  }
                }
            }
            return dp[n] = max
        }
        let res = f(n, cuts, dp)
        return  res == Number.MIN_VALUE ? 0 : res
    }
}
