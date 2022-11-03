
/**
 * @param {BigInt[]} hist
 * @param {number} n
 * @returns {BigInt}
*/
class Solution 
{
    //Function to find largest rectangular area possible in a given histogram.
    getMaxArea(hist, n)
    {
        // code here
        let stack = new Array()
        let left_bounds = new Array()
        let right_bounds = new Array()
        let maximum_area = 0
        let top_of_stack = function(stack) {
            return stack[stack.length-1]
        }
        let calculate_area = function() {
        }

        let find_left_bounds = function() {
            for(let i=0; i<n; i++) {
                if(stack.length == 0) {
                    stack.push(0)
                    left_bounds.push(0)
                } else if(hist[i] >=  hist[top_of_stack(stack)] ) {
                 
                    left_bounds.push(top_of_stack(stack) + 1)
                    stack.push(i)
                } else {
                    while(hist[i] <= hist[top_of_stack(stack)]) 
                        stack.pop()
                    if(stack.length == 0)
                        left_bounds.push(0)
                    else
                        left_bounds.push(top_of_stack(stack) + 1)
                    stack.push(i)
                }
                
            }
        }


        let find_right_bounds = function() {
            for(let i = n-1; i>=0; i--) {
                if(stack.length == 0) {
                    stack.push(n-1)
                    right_bounds.push(n-1)
                } else if(hist[i] >=  hist[top_of_stack(stack)] ) {
                 
                    right_bounds.push(top_of_stack(stack) - 1)
                    stack.push(i)
                } else {
                    while(hist[i] <=  hist[top_of_stack(stack)]) 
                        stack.pop()
                    if(stack.length == 0)
                        right_bounds.push(n-1)
                    else
                        right_bounds.push(top_of_stack(stack) - 1)
                    stack.push(i)
                }
                
                
                

            }
        }

        let find_maximum_area = function() {
            let current_area
            for(let i = 0; i<n; i++) {
                current_area = Number(hist[i]) * Number((right_bounds[i] - left_bounds[i] + 1))
                if(current_area > maximum_area)
                    maximum_area = current_area
            }
        }
        

        find_left_bounds()
        //console.log(left_bounds)
        //clear
        while(stack.length != 0)
             stack.pop()
        
        find_right_bounds()
        right_bounds.reverse()
        //console.log(right_bounds)
        find_maximum_area()

        return maximum_area
        
        
    }

    
    
}


let s = new Solution()
let hist = [9, 10, 4, 10, 4, 5, 16]
let n = hist.length
console.log(s.getMaxArea(hist, n))
