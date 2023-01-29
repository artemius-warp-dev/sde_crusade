//User function Template for javascript


/**
 * @param {string} str
 * @returns {string}
*/

class Solution{
    removeDups(str){
        //code here
        let s = new Set()
        let result = ""

        for(let i = 0; i<str.length; i++) {
            s.add(str[i])
        }

        s.forEach(value => result = result + value)
        return result
    }
}


let solution = new Solution()
let str = "zvvo"
console.log(solution.removeDups(str))
