//User function Template for javascript

/**
 * @param {string} str
 * @param {string} patt
 * @returns {number}
*/
class Solution {
    minIndexChar(str, patt){
        // code here
        let hash_map = new Map()
        let min_index = Number.MAX_VALUE
        for (let i =0; i<str.length; i++) {
            if(!hash_map.has(str[i]))
                hash_map.set(str[i], i)
               
        }

       // console.log(hash_map)
        
        for(let i = 0; i<patt.length; i++) {
            if(hash_map.has(patt[i]) && hash_map.get(patt[i]) < min_index) {
                min_index = hash_map.get(patt[i])
            }
        }

        //console.log(min_index)

        if (min_index < Number.MAX_VALUE)
            return min_index
        else
            return -1
    }
}

let solution = new Solution()
let str = "adcffaet"
let patt = "onkl"

console.log(solution.minIndexChar(str, patt))
