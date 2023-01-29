//User function Template for javascript

/**
 * @param {number} N
 * @return {number[]}
 */

class Solution {
  
  minPartition(N) {
    //code here
    let currency =  [1, 2, 5, 10, 20, 50, 100, 200, 500, 2000]
    let ans = []
    for(let i= currency.length - 1; i>=0; i--) {
        while(N >= currency[i]) {
            ans.push(currency[i])
            N-=currency[i]
        }
    }
    return ans
  }
}
