//User function Template for javascript

/**
 * @param {Number} N
 * @returns {Number}
*/
// Function to calculate the longest consecutive ones
class Solution {
    
    maxConsecutiveOnes(N)
    {
        // code here
        let count = 0
        while(N != 0) {
            N = N & (N << 1)

            count+=1
        }

        return count
    }
}

let s = new Solution()
console.log(s.maxConsecutiveOnes(14))
