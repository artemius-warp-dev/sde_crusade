//User function Template for javascript

 /**
 * @param {string} S
 * @return {string[]}
 */

class Solution {
    result = new Array()

    find_permutation_util(s, ans) {
        if(s.length == 0) {
            this.result.push(ans)
        }

        for(let i =0; i<s.length; i++) {
            let ch = s[i]
            let left = s.slice(0, i)
            let right = s.slice(i+1)
            let rest = left + right
            this.find_permutation_util(rest, ans + ch)
        }
    }

    find_permutation(S){
        //code here
        this.find_permutation_util(S, "")
        let set = new Set(this.result.sort())
        return Array.from(set)
 
    }
}


let s = new Solution()
let str = "ABB"
console.log(s.find_permutation(str))

//console.log("BC".slice(0,2))
