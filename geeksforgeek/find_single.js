
//User function Template for javascript


/**
 * @param {number} n
 * @param {number[]} arr
 * @returns {number}
*/

class Solution{
    findSingle(n, arr){
        //code here
        let res = 0
        for (let e of arr) {
            //console.log([e,res])
            res = res ^ e
            //console.log(res)
        }
        return res
    }
}


let s = new Solution()
let arr = [1, 2, 3, 5, 3, 2, 1, 4, 5, 6, 6]
console.log(s.findSingle(arr.length, arr))
