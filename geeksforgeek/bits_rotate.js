
//User function Template for javascript



/**
 * @param {number} n
 * @param {number} d
 * @returns {number[]}
*/


//10101
//2
//left -> 10100 = (N << 2) | (N >> 5 - 2)

class Solution{
    rotate(n,d){
        //code here
        var BITS = 16

        d = d % BITS
        let left_rotate = function(n,d) {
            return (n << d | n >> (BITS - d)) & 0xFFFF
        }

        let right_rotate = function(n,d) {
            return (n >> d | n << (BITS - d)) & 0xFFFF
        }

        return [left_rotate(n,d), right_rotate(n,d)]

    }
}



let s = new Solution()
console.log(s.rotate(321, 18))
