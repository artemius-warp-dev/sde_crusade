//User function Template for javascript
/**
 * @param {string} s
 * @returns {string}
*/
 

class Solution 
{

    reverse(word) {
        let char_arr = word.split("")
        let start = 0, end = char_arr.length - 1
        let tmp
        while (start <= end) {
            tmp = char_arr[start]
            char_arr[start] = char_arr[end]
            char_arr[end] = tmp
            start = start + 1
            end = end - 1
        }
        
        return char_arr.join("")
    }

    //Function to reverse words in a given string.
    reverseWords(s)
    {
        // code here
        let list = s.split(".")
        //console.log(list)
        let final_str = ""
        for(let i = 0; i < list.length-1; i++) {
            final_str = final_str  + this.reverse(list[i]) + "."
        }

        final_str = final_str + this.reverse(list[list.length - 1])
        return this.reverse(final_str)
    }
}



let str = "i.like.this.program.very.much"
let solution = new Solution()
console.log(solution.reverseWords(str))
