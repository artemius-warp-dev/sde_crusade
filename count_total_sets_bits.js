//User function Template for javascript

/**
 * @param {Number} N
 * @returns {Number}
*/

function max2(n) {
    let c = 0 

    while((1<<c) <= n) {
        c+=1
    }
    return c - 1
}


class Solution {
    //Function to return sum of count of set bits in the integers from 1 to n.
    countSetBits(N)
    {
        // code here
        //console.log(N)
        if(N == 0) {
            return 0
        }

        let x = max2(N)
        let remaining_part = (N - (1 << x))
        // x * 2^(n-1) + (n - (1<<x)  remainder + rec
        let ans = x * (1 << (x-1)) + remaining_part + 1 + this.countSetBits(remaining_part)  
        return ans
        
    }
}


let s = new Solution()
console.log(s.countSetBits(4))
