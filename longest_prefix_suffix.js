//User function Template for javascript

/**
 * @param {string} s
 * @returns {number}
 */

class Solution {
    lps(s){ 
        //code here
        let n = s.length
        let lps = new Array(n).fill(0)
        
        lps[0] = 0
        
        let len = 0
        
        let i = 1
        while(i < n) {
            //console.log(s.charAt(len), s.charAt(i), len)
            if(s.charAt(i) == s.charAt(len)) {
                len++
                lps[i] = len
                i++
            } else {
                if(len != 0) {
                    //console.log('else')
                    len = lps[len-1]
                    //console.log(len)
                    
                } else {
                    lps[i] = 0
                    i++
                }
            }
        }
        
        console.log(lps)
        let res = lps[n-1]
        
        //let half = Math.floor(n/2)
        return  res //(res > half) ? half : \res
        
    }
}
