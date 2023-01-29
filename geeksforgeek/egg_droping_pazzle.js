//User function Template for javascript


/**
 * @param {number} n
 * @param {number} k
 * @returns {number}
*/

class Solution 
{
    //Function to find minimum number of attempts needed in 
    //order to find the critical floor.
    eggDrop(n, k)
    { 
        // code here
        let dp = new Array(n+1).fill().map(() => new Array(k+1).fill(-1))
        for(let i=1; i<=n; i++) {
            dp[i][0] = 0
            dp[i][1] = 1 
        }
        
        for(let i=1; i<=k; i++) {
            dp[1][i] = i
        }
        //console.log(dp)
        for(let i=2; i<=n; i++) {
            for(let j=2; j<=k; j++) {
                dp[i][j] = Number.MAX_VALUE
                for(let x = 1; x<=j; x++) {
                    let res = 1 + Math.max(dp[i-1][x-1], dp[i][j-x]) //brake and not brake cases
                    if(res < dp[i][j]) {
                        dp[i][j] = res
                    }
                }
            }
        }
        
        return dp[n][k]
        
        
    }
}

