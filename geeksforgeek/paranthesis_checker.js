//User function Template for javascript

/**
 * @param {string} x
 * @returns {boolean}
*/

class Solution
{
    //Function to check if brackets are balanced or not.
    ispar(x)
    {
        //your code here
        let stack = []
        for(let i = 0; i<x.length; i++) {
            if(x[i] == "{" || x[i] == "(" || x[i] == "[") {
                stack.push(x[i])
                continue
            }

            //console.log(stack)
            if(stack.length == 0)
                return false


            let check
            //console.log(x[i])
            switch(x[i]) {
            case '}':
                check = stack.pop()
                if(check == "[" || check == "(")
                    return false
                break
            case ')':
                check = stack.pop()
                if(check == "[" || check == "{")
                    return false
                break
            case ']':
                check = stack.pop()
                if(check == "{" || check == "(")
                    return false
                break
            }
           

        }
        //console.log(stack)
        return stack.length == 0
        

        
    }
}

let s = new Solution()
console.log(s.ispar("{([])}"))
