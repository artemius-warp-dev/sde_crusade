//User function Template for javascript


/**
 * @param {number} x
 * @param {number} y
 * @param {string} s1
 * @param {string} s2
 * @returns {number}
*/
 
class Solution
{
    //Function to find the length of longest common subsequence in two strings.
    lcs(x, y, s1, s2)
    {
        // code here
        
        let cache = []
        for(let i=0; i<x+1; i++) {
            cache[i] = new Array(y+1).fill(-1)
        }
        
        
        let f = function(i,j, s1,s2) {
            if(i == 0  || j == 0) return 0
            
            if(cache[i][j] != -1) return cache[i][j]
            //console.log(s1[i], s2[i])
            if(s1[i-1] === s2[j-1]) {
                cache[i][j] = 1 + f(i-1, j-1, s1, s2)
                
            } else {
                cache[i][j] =  Math.max(f(i-1, j, s1, s2), f(i,j-1, s1,s2))
            }
            //console.log(s1[i-1], s2[j-1], cache[i][j])
            return cache[i][j]
        }
        
        return f(x,y,s1,s2)
    }
}
