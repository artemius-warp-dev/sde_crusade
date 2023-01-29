//User function Template for javascript

/**
 * @param {number[]} arr
 * @param {number} n
 * @returns {number}
*/

const INT_BITS = 32

class Solution {
    
    maxSubsetXOR(arr, n)
    {
        // code here
        let index = 0
        
        for(let i=INT_BITS - 1; i >= 0; i--) {
            let maxInd = index
            let maxEle = Number.MIN_VALUE
            
            for(let j=index; j<n; j++) {
                
                if( (arr[j] & (1 << i)) !=0 && arr[j] > maxEle) {
                    maxEle = arr[j]
                    maxInd = j
                }
            }
            
            if(maxEle == Number.MIN_VALUE) {
                continue
            }
            
            
            let tmp = arr[index]
            arr[index] = arr[maxInd]
            arr[maxInd] = tmp
            
            maxInd = index
            
            for(let j =0; j<n; j++) {
                
                if(j != maxInd && (arr[j] & (1 << i)) != 0)
                    arr[j] = arr[j] ^ arr[maxInd]
            
                
            }
            
            index++
            
        }
        
        let res = 0
        for(let i=0; i<n; i++) {
            res ^= arr[i]
        }
        return res
    }
}
