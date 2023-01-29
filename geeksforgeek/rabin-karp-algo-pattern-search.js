//User function Template for javascript

/**
 * @param {string} pat
 * @param {string} txt
 * @returns {number[]}
*/



class Solution {
    
    search(pat, txt)
    {
        //your code here
        let m = pat.length
        let n = txt.length
        let q = 101
        let alphabet = 256
        
        let i
        let j
        
        let p = 0
        let t = 0
        let h = 1
        
        for(let i=0; i < m-1; i++) {
            h = (h * alphabet) % q
        }
        
        
        for(i=0; i<m; i++) {
            p = (alphabet * p + pat[i].charCodeAt()) % q
            t = (alphabet * t + txt[i].charCodeAt()) % q 
        }
        
        let res = []
        for(i=0; i <= n - m; i++) {
            if(p == t) {
                for(j = 0; j < m; j++) {
                    if(txt[i+j] != pat[j])
                        break   
                }
                
              if(j == m) {
                  //console.log(i,j)
                  res.push(i + 1)
              }  
            }
            
            if( i < n - m) {
                t = (alphabet * (t - txt[i].charCodeAt() * h) + txt[i + m].charCodeAt()) % q
                
                if(t < 0)
                    t = (t + q)
            }
            
        }
        
        //console.log(res.length)
        if (res.length > 0) {
            return res
        } else {
            return [-1]
        }
    }
}
