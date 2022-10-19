//User function Template for javascript

/**
 * @param {Number} n
 * @returns {boolean}
*/


class Solution 
{
    //Function to check if the number is sparse or not.
    isSparse(n)
    {
        // code here
        if(n & (n >> 1)) {
            return 0
        }

        return 1
    }
}

let s = new Solution()
console.log(s.isSparse(4))
