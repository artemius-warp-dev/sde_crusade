/**
 * @param {string} s
 * @return {number}
 */

class Solution {
    atoi(s){
        //code here

        let n = s.length
        let sign = 1
        let base = 0
        
        for (let i=0; i < n; i++) {
            if (s.charAt(i) == " ") {
                continue
            }
            //console.log(i)

            if (s.charAt(i) == '-') {
                sign = -1
                i = i+1
            }
            if (s[i] >= '0' && s[i] <= '9' && s[i+1] != "-") {

                base = base * 10 + s[i].charCodeAt(0) - '0'.charCodeAt(0)
            } else {
                return -1
            }
        }

        return base * sign
        
    }
}



let solution = new Solution()
let str = "--123"
console.log(solution.atoi(str))
