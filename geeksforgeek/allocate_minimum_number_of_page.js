//User function Template for javascript

/**
 * @param {number[]} a
 * @param {number} n
 * @param {number} m
 * @returns {number}
*/

class Solution
{
    
    is_posible(arr, pages, students) {
        let cnt = 0
        let sumAllocated = 0
        for(let i=0; i<arr.length; i++) {
            if(sumAllocated + arr[i] > pages) {
                cnt++
                sumAllocated = arr[i]
                if(sumAllocated > pages)  return false
                
            } else {
                sumAllocated += arr[i]
            }
        }
        if(cnt < students) return true
        return false
    }
    
    //Function to find minimum number of pages.
    findPages(a, n, m)
    {
        //your code here
        if(m > n) return -1
        let low = a[0]
        let high = 0
        
        for(let i=0; i<n; i++) {
            high = high + a[i]
            low = Math.min(low, a[i])
        }
        
        while(low <= high) {
            let mid = Math.floor((low + high)/2)
            if(this.is_posible(a,mid,m)) {
                high = mid - 1
            } else {
                low = mid + 1
            }
        }
        
        return low
        
        
    }
}
