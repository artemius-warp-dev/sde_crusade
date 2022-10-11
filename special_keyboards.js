//User function Template for javascript

/**
 * @param {number} N
 * @returns {number}
*/

class Solution {
    
    rec_optimalKeys(N) {

        if(N<7) return N
        let multiplier = 2
        let max = 0
        for(let i = N-3; i>=1; i--) {
            let cur = this.optimalKeysUtil(i) * multiplier
            if(cur > max) {
                max = cur
            }
            
           multiplier+=1
        }
         return max
    }


    cash_optimalKeys(N) {
        let cash = new Array(N)
        for(let i = 1; i<=6; i++) {
            cash[i-1] = i
        }

        for(let n = 7; n<=N; n++) {
            cash[n-1] = 0
            let multiplier = 2
            for(let b = n-3; b>=1; b--) {
                let cur = cash[b-1] * multiplier
                //console.log(cur)
                if(cur > cash[n-1]) {
                    cash[n-1] = cur
                }
                
                multiplier+=1
            }
        
        }
        return cash[N-1]
    }

    optimalKeys(N)
    {
        //your code here
        return this.cash_optimalKeys(N)
    }
}

let s = new Solution()
let N = 11
console.log(s.optimalKeys(N))
