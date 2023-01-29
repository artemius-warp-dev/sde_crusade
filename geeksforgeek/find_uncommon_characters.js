//User function Template for javascript


/**
 * @param {string} A
 * @param {sring} B
 * @returns {string}
*/

class Solution{
    UncommonChars(A, B){
        //code here
        let result = ""
        let hash_map = new Map()
        for (let i = 0; i<A.length; i++) {

            hash_map.set(A[i], 1)

            
        }


        for(let i = 0; i<B.length; i++) {
            if(hash_map.has(B[i])) {
                hash_map.set(B[i], -1)
            } else {
                hash_map.set(B[i], 1)
            }

        }

        hash_map.forEach((value, key) => {
            if(value != -1)
                result = result + key
        })
        return [...result].sort().join("")
    }
}


let solution = new Solution()
let str1 = "geeksforgeeks"
let str2 = "geeksquiz"
console.log(solution.UncommonChars(str1, str2))
