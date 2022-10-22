//User function Template for javascript


/**
 * @param {number[]} price
 * @param {number} n
 * @returns {number[]}
*/
class Solution
{
    //Function to calculate the span of stockâ€™s price for all n days.
    calculateSpan(price, n)
    {
        // code here
        let stack = []
        let span = []
        for(let i =0; i<n; i++) {
            //console.log([price[i],  price[stack.length - 1]])
            //console.log(stack)
            while(stack.length != 0 && price[stack[stack.length - 1]] <= price[i]) {
               // console.log(parice[i])
                stack.pop()
            }

            if(stack.length == 0) {
                span.push(i + 1)
            } else {
                //console.log(stack[stack.length - 1])
                span.push(i - stack[stack.length -1])
            }
            

            stack.push(i)
            //console.log(stack.length)

 
        }
        return span
    }
}

let s = new Solution()
let N = 3
let price = [68, 735, 101]
console.log(s.calculateSpan(price, N))
