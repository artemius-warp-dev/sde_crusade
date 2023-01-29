//User function Template for javascript

/**
 * @param {Number} n
 * @returns {Number}
*/


class Solution 
{
    //Function to find position of first set bit in the given number.
    getFirstSetBit(n)
    {
        // code here
        if(n==0) return 0
        let position = 1
        let m = 1
        while((n & m) == 0) {
            m = m << 1
            if(position >= Math.pow(n,10)) {
                return -1
            }
            position+=1
        }
        return position
        
    }
}


let s = new Solution()
let n = 12
console.log(s.getFirstSetBit(n))
