//User function Template for javascript


/**
 * @param {number} W
 * @param {number[]} wt
 * @param {number[]} val
 * @param {number} n
 * @returns {number}
*/

class Solution 
{
    //Function to return max value that can be put in knapsack of capacity W.
    knapSack(W, wt, val, n)
    { 
       // code here
       let res = []
       for(let i =0; i<=n; i++) {
           res[i] = new Array()
           for(let j=0; j<=W; j++) {
               res[i][j] = -1
           }
       }
      //console.log(res)
      let calculate = function(n, w) {
        //console.log([n, w])   
        let result    
        if(res[n][w] != -1) return res[n][w]
           
         if(n == 0 || w == 0) {
            result = 0
        } else if(wt[n-1] > w) {
            result = calculate(n-1, w)
        } else {
            let tmp1 = calculate(n-1, w)
            let tmp2 = val[n-1] + calculate(n-1, w - wt[n-1])
            result = Math.max(tmp1, tmp2)
        }
      // console.log(result)
        res[n][w] = result
    
        return result
      }
       
       
      return calculate(n, W)
       
     }
}

