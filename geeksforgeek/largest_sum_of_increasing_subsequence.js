//User function Template for javascript

/**
 * @param {number[]} arr
 * @param {number} n
 * @return {number}
 */

class Solution {
    maxSumIS(arr,n){
        //code here
        let sum = 0
        let dp = new Array(n).fill(new Array(n+1).fill(-1))
        //console.log(dp)
        let f = function(ind, prev, arr, n, dp) {
            if(ind >= n) return 0
            
            if(dp[ind][prev + 1] != -1) return dp[ind][prev+1]
            
            let take = 0
            let no_take = f(ind+1, prev, arr, n, dp)
            if(prev == -1 || arr[ind] > arr[prev]) {
                take = arr[ind] + f(ind+1, ind, arr, n, dp)
            }
            //console.log([take, no_take])
            return dp[ind][prev+1] = Math.max(take, no_take)
        }
        
       
        
        return  f(0,-1, arr, n, dp)
        
   
    }
}
