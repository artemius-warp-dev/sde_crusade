//User function Template for javascript

/**
 * @param {string} a
 * @return {string}
 */
 
 String.prototype.replaceAt = function(index, replacement) {
     return this.substring(0, index) + replacement + this.substring(index + replacement.length);
 }
 
 
class Solution {
  chooseandswap(a) {
    //code here
    let alphabet_size = 26
    let n = a.length
    let cache = new Array(alphabet_size).fill(-1)
    let i,j
    for(i=0; i<n; i++) {
        let index = a[i].charCodeAt(0) - 'a'.charCodeAt(0)
        if(cache[index] == -1) {
            cache[index] = i
        }
    }
    
    for(i=0; i<n; i++) {
        let flag = false
        for(j = 0; j < a[i].charCodeAt(0) - 'a'.charCodeAt(0); j++) {
            if(cache[j] > cache[a[i].charCodeAt(0) - 'a'.charCodeAt(0)]) {
                flag = true
                break
            }
        }
        if(flag) break
    }
    
    if(i < n-1) {
        let ch1 = a[i]
        let ch2 = String.fromCharCode(j + 'a'.charCodeAt(0));
        //console.log(ch1, ch2)
        for(i=0; i<n; i++) {
            if(a[i] == ch1){
               a=a.replaceAt(i, ch2)
            } else if(a[i] == ch2) {
               a=a.replaceAt(i,ch1)
            }
        }
    }
    
    return a
    
  }
}
