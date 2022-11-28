/**
 * @param {number} n
 * @param {number} s
 * @returns {string}
*/

class Solution{
    findLargest(n,s){
        //code here
        if(s ==0) return n == 1 ?  0 : -1
        if(s > n * 9) return -1
        let res = []
        for(let i=0; i<n; i++) {
            if(s >= 9) {
                res.push(9)
                s-=9
            } else {
                res.push(s)
                s = 0
                //break
            }
        }
       // console.log(res)
        let number = ""
        for(let i=0; i< res.length; i++) {
            number+=res[i]
        }
        return number
        
    }
}


