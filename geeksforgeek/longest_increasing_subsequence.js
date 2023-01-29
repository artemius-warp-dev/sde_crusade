//User function Template for javascript

/**
 * @param {number} n
 * @param {number[]} a
 * @returns {number}
*/
class Solution 
{
    //Function to find length of longest increasing subsequence.
    longestSubsequence(n, a)
    {
        // code here
        let ans = []
        ans.push(a[0])
        
        let lower_bound = function(list, x, l,r) {
            if (l > r) return l
            let mid = Math.floor(l + (r-l)/2)
            if(list[mid] >= x) {
                return lower_bound(list, x, l, mid - 1)
            } else {
                return lower_bound(list, x, mid + 1, r)
            }
        }
        
        
        for(let i=1; i<n; i++) {
           if(a[i] > ans[ans.length - 1]) {
               ans.push(a[i])
           }  else {
               
               let index = lower_bound(ans, a[i], 0, ans.length  - 1) 
               //console.log(["!" ,a[i], ans[ans.length - 1], index])
               ans[index] = a[i]
               //console.log(ans)
           }
        }
        
        return ans.length
    }
}
