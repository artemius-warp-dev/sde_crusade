//User function Template for javascript

/**
 * @param {number} N
 * @param {number} from
 * @param {number} to
 * @param {number} aux
 * @returns {number}
*/
class Solution {
    count = 0
    toh(N, from, to, aux)
    {
        // code here
        if(N == 0) return N

        this.toh(N-1, from, aux, to)
        console.log("move disk", N, "from rod", from, "to rod", to)
        this.count+=1
        this.toh(N-1, aux, to, from)
        //console.log(this.count)
        return this.count
        
    }
}



let s = new Solution()
let N = 3
s.toh(3, 'A', 'C', 'B')
    
    
