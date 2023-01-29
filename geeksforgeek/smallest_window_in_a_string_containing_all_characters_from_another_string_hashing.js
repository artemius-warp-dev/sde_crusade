//User function Template for javascript

/**
 * @param {string} s
 * @param {string} p
 * @returns {string}
*/

const no_of_chars = 256


class Solution 
{
    //Function to find the smallest window in the string s consisting
    //of all the characters of string p.
    smallestWindow(s, p)
    {
        // code here
        let len1 = s.length
        let len2 = p.length
        
        if(len1 < len2) {
            return -1
        }
        
        let hash_str = new Array(no_of_chars).fill(0)
        let hash_ptr = new Array(no_of_chars).fill(0)
        
        for(let i = 0; i<len2; i++) {
            hash_ptr[p.charAt(i).charCodeAt(0)]++
        }
        
        let start = 0, start_index = -1
        let min_len = Number.MAX_VALUE
        
        let count = 0
        for(let j=0; j<len1; j++) {
            hash_str[s.charAt(j).charCodeAt(0)]++
            
            
            //console.log(hash_str[s.charAt(j).charCodeAt(0)] , hash_ptr[s.charAt(j).charCodeAt(0)])
            if(hash_str[s.charAt(j).charCodeAt(0)] <= hash_ptr[s.charAt(j).charCodeAt(0)])
                count++
            //console.log(count)    
            
            if(count == len2) {
                //console.log(count)
                while(hash_str[s.charAt(start).charCodeAt(0)] > hash_ptr[s.charAt(start).charCodeAt(0)] 
                || hash_ptr[s.charAt(start).charCodeAt(0) == 0]) {
                    
                    if(hash_str[s.charAt(start).charCodeAt(0)] > hash_ptr[s.charAt(start).charCodeAt(0)])
                        hash_str[s.charAt(start).charCodeAt(0)]--
                    
                    start++
                    
                }
                
                let len_window  = j - start + 1
                if(min_len > len_window) {
                    min_len = len_window
                    start_index = start
                }
            }
        }
        
        if(start_index == -1) 
            return -1
        let res = s.substring(start_index, start_index + min_len)
        //console.log(res)
        return res
    }
}
