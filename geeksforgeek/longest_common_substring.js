//User function Template for javascript

/**
 * @param {string} S1
 * @param {string} S2
 * @param {number} n
 * @param {number} m
 * @return {number}
 */

class Solution {
    longestCommonSubstr(S1,S2,n,m){
        //code here
        let cache = []
        for(let i=0; i<n+1; i++) {
            cache[i] = new Array(m+1).fill(0)
        }
        
        let ans = 0
        for(let i=1; i<=n; i++) {
            for(let j=1; j<=m; j++) {
                 if(S1[i-1] == S2[j-1]) {
                    cache[i][j] = 1 + cache[i-1][j-1]
                    ans = Math.max(ans, cache[i][j])
                 } else {
                     cache[i][j] = 0
                 }
            }
        }
        
        
        return ans
    }
}
