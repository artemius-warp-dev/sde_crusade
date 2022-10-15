//User function Template for javascript
/**
 * @param {Number} n
 * @param {Number} k
 * @returns {boolean}
*/

class Solution 
{
    // Function to check if Kth bit is set or not.
    checkKthBit(n, k)
    {
        //code here
        
        let temp = n >> (k)
        //console.log(temp.toString(2))
        //console.log(n.toString(2))
        return (temp & 1) > 0
    }
}

let s = new Solution()
let n = 500, k = 3
console.log(s.checkKthBit(n,k))
