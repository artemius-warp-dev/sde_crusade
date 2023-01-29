//User function Template for javascript
/**
 * @param {string} s
 * @param {string} x
 * @returns {number}
*/
 
class Solution 
{
    //Function to locate the occurrence of the string x in the string s.
    strstr(s, x)
    {
        // code here
        let N = s.length, M = x.length
        for(let i=0; i< N-M; i++) {
            for(let j=0; j<M; j++) {
                if(s[i+j] != x[j]) {
                    break
                }
                if (j == M - 1) 
                    return i
                
            }
        }
        return -1
    }
}


let solution = new Solution()
let s1 = "GeeksForGeeks", s2 = "For"
console.log(solution.strstr(s1, s2))
