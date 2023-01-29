// User function Template for javascript

/**
 * @param {string} s
 * @param {string} t
 * @return {number}
 */

class Solution {
    editDistance(s, t) {
        // code here
        let n = s.length, m=t.length
        // let dp = new Array(n)
        // for(let i = 0; i<n; i++) {
        //     dp[i] = new Array(m).fill(-1)
        // }
        let dp = new Array(n).fill().map(() => new Array(m).fill(-1))
        let f = function(i,j, s,t, dp) {
           // console.log([s[i], t[j]])
            if(i < 0) return j+1
            if(j < 0) return i+1
            //if(j == 0) return i
            //if(i == 0) return j
            if(dp[i][j] != -1) return dp[i][j]
            if(s[i] == t[j]) return dp[i][j] = f(i-1, j-1, s,t, dp)
            let r = 1 + Math.min(f(i-1, j, s,t, dp), Math.min(f(i-1, j-1, s, t, dp),  f(i, j-1, s, t, dp)))
            //console.log(r)
            return dp[i][j] = r
        }
        
        return f(n-1,m-1, s, t, dp)
    }
}
