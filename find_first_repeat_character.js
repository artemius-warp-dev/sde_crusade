//User function Template for javascript

/**
 * @param {string} s
 * @return {string}
 */

class Solution {
  firstRepChar(s) {
    // code here
    let cache = new Map()
    for(let ch of s) {
        if(cache.has(ch)) {
            return ch
        } else {
            cache.set(ch, 1)
        }
    }
    return -1
  }
}
