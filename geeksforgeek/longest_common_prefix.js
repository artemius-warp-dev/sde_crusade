//User function Template for javascript

/**
 * @param {string[]} arr
 * @param {number} n
 * @returns {string}
 */

class Solution {


   longestCommonPrefixUtil(s1, s2) {
        let n1 = s1.length, n2 = s2.length
        let i = 0, j = 0
        let result = ""
       while(i < n1 && j < n2) {
           if (s1[i] != s2[j]) {
                break
            }
           //console.log(s1[i])
           result = result + s1[i]
           i=i+1
           j=j+1
        }
        return result
    }

    
    longestCommonPrefix(arr,n){ 
        //code here
        let prefix = arr[0]
        for (let i = 1; i<n; i++) {
            prefix = this.longestCommonPrefixUtil(prefix, arr[i])
        }
        if (prefix == "")
            return -1
        else
            return prefix
    }
}

let solution = new Solution()
let arr = ["d", "d", "d", "d"]
let N = arr.length
console.log(solution.longestCommonPrefix(arr, N))
