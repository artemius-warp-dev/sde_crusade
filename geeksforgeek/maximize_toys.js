//User function Template for javascript

/**
 * @param {number[]} arr
 * @param {number} N
 * @param {number} K
 * @returns {number}
*/

class Solution {
    
    toyCount(arr, N, K)
    {
        //your code here
        let count = 0
        let sum = 0
        arr.sort((a,b) => a-b)
       // console.log(arr)
        let i = 0
        while(arr[i] + sum <= K) {
            sum+=arr[i]
            count++
            i++
        }
        return count
    }
}
