//User function Template for javascript

/**
 * @param {Number} m
 * @param {Number} n
 * @returns {Number}
*/
class Solution 
{
    //Function to find the first position with different bits.
    posOfRightMostDiffBit(m, n)
    {

        let rightmost_pos = function(n) {
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


        // code here
        if(m==n) return -1
        
        return rightmost_pos(m^n)


        
    }
}

let s = new Solution()
let n = 52
let m = 4
console.log(s.posOfRightMostDiffBit(m,n))
