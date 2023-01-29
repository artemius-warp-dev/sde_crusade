/**
 * @param {string} a
 * @param {string} b
 * @returns {boolean}
*/
 
class Solution 
{
    //Function is to check whether two strings are anagram of each other or not.
    isAnagram(a, b)
    {
        // code here
        if(a.length != b.length) {
            return false
        }

        let map = new Map()

        for(let i =0; i<a.length; i++) {
            if(map.has(a[i])) {
                map.set(a[i], map.get(a[i]) + 1)
            } else {
                map.set(a[i], 1)
            }
        }


        for (let i=0; i<b.length; i++) {
            if(map.has(b[i])) {
                map.set(b[i], map.get(b[i]) - 1)
            } else {
                return false
            }
        }

        let keys = map.keys()
        for(let key of keys) {
            if(map.get(key) != 0) {
                return false
            }
        }

        return true

        

        
        
    }
}


let solution = new Solution()
let a = "allergy"
let b = "allergic"

console.log(solution.isAnagram(a,b))
