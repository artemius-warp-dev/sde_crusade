//User function Template for javascript


/**
 * @param {number} n
 * @returns {BigInt}
*/

class Solution 
{
    //Function to count the number of ways in which frog can reach the top.
    countWays(n)
    {
        // code here
        let res = []
        res[0] = 1
        res[1] = 1
        res[2] = 2
        res[3] = 4
        for(let i = 3; i<=n; i++) {
            res[i] = (res[i-3] + res[i-2] + res[i-1]) % (Math.pow(10, 9) + 7)
        }
       // console.log(res[n])
        return res[n]
        
    }
}
