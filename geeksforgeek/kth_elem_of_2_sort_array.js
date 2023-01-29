//User function Template for javascript

/**
 * @param {number[]} A
 * @param {number[]} B
 * @param {number} n
 * @param {number} m
 * @param {number} k
 * @returns {number}
 */

class Solution {
    kthElement(A,B,n,m,k){ 
        //code here
        if(n > m) {
            return this.kthElement(B, A, m, n, k)
        }
        
        let low = Math.max(0, k-m)
        let high = Math.min(k,n)
        
        while(low <= high) {
            let cut1 = Math.floor((low + high) / 2)
            let cut2 = k - cut1
            let l1 = cut1 == 0 ? Number.MIN_VALUE : A[cut1 - 1]
            let l2 = cut2 == 0 ? Number.MIN_VALUE : B[cut2 - 1]
            let r1 = cut1 == n ? Number.MAX_VALUE : A[cut1]
            let r2 = cut2 == m ? Number.MAX_VALUE : B[cut2]
            //console.log(l1,r1,l2,r2)
            if(l1 <= r2 && l2 <= r1) {
                return Math.max(l1,l2)
            } else if(l1 > r2) {
                high = cut1 - 1
            } else {
                low = cut1 + 1
            }
        }
    }
}
