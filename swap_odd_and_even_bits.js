//User function Template for javascript

/**
 * @param {Number} n
 * @returns {Number}
*/

class Solution
{
    //Function to swap odd and even bits.
    swapBits(n)
    {
        // code here
        let even_bits = n & 0xAAAAA
        let odd_bits = n & 0x55555

        even_bits >>= 1
        odd_bits <<= 1

        return even_bits | odd_bits
    }
}

let s = new Solution()
console.log(s.swapBits(23))
