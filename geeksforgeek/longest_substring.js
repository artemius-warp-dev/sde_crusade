//User function Template for javascript

/**
 * @param {string} s
 * @return {number}
 */
class Solution {
    
  longestSubstrDistinctChars(s) {
      //code here
      
      let start_index = 0
      let seen = new Map()
      let result = 0
      for (let i = 0; i < s.length; i++) {

          if (seen.has(s[i])) {
              start_index = Math.max(start_index, seen.get(s[i]) + 1)
          }
          
          
          result = Math.max(result, i - start_index + 1)
          seen.set(s[i], i)
      }

      return result

      
  }
}

let solution = new Solution()
let str = "aaa"
console.log(solution.longestSubstrDistinctChars(str))
