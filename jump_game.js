//User function Template for javascript

/**
 * @param {number[]} arr
 * @param {number} n
 * @return {number}
 */

class Solution {
     minJumps(arr,n){
         //code here
    // TLE recursion memoization version
    //     let dp = new Array(n).fill(-1)
    //     let f = function(i, arr, dp) {
    //         if(i >= arr.length) return Number.MAX_VALUE
    //         if(i == arr.length - 1) return 0
    //         if (dp[i] != -1) return dp[i]
    //         let res = Number.MAX_VALUE
    //         for(let j = 1; j<=arr[i]; j++) {
    //             res = Math.min(res, 1 + f(i +j, arr, dp))
    //         }
           
    //         return dp[i] = res

    //     }
    //     let r = f(0, arr, dp)
    //     if(r == Number.MAX_VALUE) return -1
    //     else return r
    
   
    let jumps = 0
    let l=0, r=0
    
    if(n<=0) return -1
    let count = 0
    while (r < n-1) {
        let farthest = 0
        for(let i=l; i<=r; i++) {
            farthest = Math.max(farthest, i + arr[i])
        }
        if(farthest == 0) return -1
        l = r + 1
        r = farthest
        jumps++
        count++
    }
    return jumps
    
    }
}
