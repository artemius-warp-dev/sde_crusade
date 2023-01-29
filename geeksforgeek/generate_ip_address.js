//User function Template for javascript

/**
 * @param {string} s
 * @returns {string[]} 
 */

var count = 0

class Solution {

    isValid(str) {
        if(str.length > 1 && str[0] == '0') return false
        if(parseInt(str) <= 255) return true
        return false
    }


    backtrack(s, idx, res, temp, seg, l) {
        if((seg == 0 && idx < l) || (seg != 0 && idx > l)) return
        // console.log("backtrack")
        //console.log(idx)
        //console.log([s, idx, res, temp, seg, l])

       if(seg === 0 && idx === l) {
           // console.log(temp)
           //console.log("hello")
           res.push(temp.slice().join('.'))
           return
       }


        for(let i=1; i<4; i++) { 
            if(idx+i>l) break
            let chunk = s.substring(idx, idx+i) // 0 to 3 character's substring
            //console.log(idx)
           if(this.isValid(chunk)) {
               temp.push(chunk)
               //count = count + 1
               // console.log(count)
               //console.log([temp, idx, seg, i])
               this.backtrack(s, idx+i, res, temp, seg-1, l)
               temp.pop()
               //console.log(["pop", temp, idx, seg, i])
            }
        }
    }

    genIp(s){
        // code here
        let res = []
        this.backtrack(s, 0, res, [], 4, s.length)
        //console.log(res)
        if (res.length == 0)
            return -1
        else
            return res
    }

    

    
}



let solution = new Solution()
let str = "1234"
//solution.genIp(str)
console.log(solution.genIp(str))
