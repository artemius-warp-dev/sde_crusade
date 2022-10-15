//User function Template for javascript

/**
 * @param {number} N
 * @param {number} L
 * @param {number} R
 * @return {number}
*/

class Solution {
    toggleBits(N,L,R){
        //code here
        //console.log(N.toString(2))
        let left = (1 << (L - 1)) -1
        //console.log(left.toString(2))
        let right = (1 << R) -1
        //console.log(right.toString(2))
        let num = left ^ right
        //console.log(num.toString(2))
        return (num ^ N)
    }
}


let s = new Solution()
let N = 17, L=2, R=3
console.log(s.toggleBits(N,L,R))
