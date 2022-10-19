//User function Template for javascript
/**
 * @param {Number} a
 * @param {Number} b
 * @returns {Number}
*/
// Function to find number of bits needed to be flipped to convert A to B
class Solution {
    
    countBitsFlip(a, b)
    {
        // code here
        console.log(a.toString(2), b.toString(2))
        var count_flip = function(n) {
            let count = 0
            console.log(n.toString(2))
            while(n != 0) {
                console.log([n.toString(2), (n-1).toString(2)])
                n &= (n-1)
               
                count+=1
            }
            return count
        }

        return count_flip(a ^ b)
    }
}

let s = new Solution()
console.log(s.countBitsFlip(10, 20))
