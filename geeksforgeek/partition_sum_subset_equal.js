//User function Template for javascript

/**
 * @param {number[]} arr
 * @param {number} n
 * @returns {boolean}
*/

class Solution {
    
    equalPartition(arr, n)
    {
        //your code here
        
        //if( n % 2 != 0) return 0
        let sum = 0
        for(const e of arr) {
            sum+=e
            //console.log(e)
        }
        if(sum % 2) return false
        let target = sum / 2
        //console.log(target)
        let dp = new Array(arr.length+1)
        for(let i=0; i<=n; i++) {
            dp[i] = new Array(target + 1).fill(-1)
        }
        
        
        let f = function(i, target, dp) {
            if(target == 0) return true
            if(i == 0) return false
            
            if(dp[i][target] != -1) return dp[i][target]
            
            if(arr[i-1] > target) {
                dp[i][target] = f(i-1, target, dp)
            } else {
                dp[i][target] = f(i-1, target-arr[i-1], dp) || f(i-1, target, dp)
            }
            
            
            
            return dp[i][target]
        }
        
        return f(arr.length, target, dp)
        
    }
    
}
