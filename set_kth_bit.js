//User function Template for javascript


/**
 * @param {number} n
 * @param {number} k
 * @returns {number}
*/

class Solution{
    setKthBit(n, k){
        //code here
        let k_temp = 1 << k
        //console.log(n.toString(2))
        //console.log(k_temp.toString(2))
        return n | k_temp
    }
}



let s = new Solution()
let N = 10, K = 2
console.log(s.setKthBit(N, K))
