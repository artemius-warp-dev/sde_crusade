//User function Template for javascript

/**
 * @param {string} S
 * @return {string} 
*/

class Solution {
  longestPalin(S){
    //code here
    let str = ""
    let n = S.length
    let max_len = 0
    for(let i=0; i<n; i++) {
        //odd
        let l = i, r = i
        while(l >= 0 && r < n && S[l] == S[r]) {
            if(r -l + 1 > max_len) {
                max_len = r - l + 1
                str = S.substring(l,r+1)
            }
            l--
            r++
        }
        
        //even
        l=i, r = i+1
        while(l >= 0 && r < n && S[l] == S[r]) {
            if(r -l + 1 > max_len) {
                max_len = r - l + 1
                str = S.substring(l,r+1)
            }
            l--
            r++
        }
    }
    //console.log(max_len)
    return str
    
  
  }
}
