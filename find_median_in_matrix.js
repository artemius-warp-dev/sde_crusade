2
//User function Template for javascript


/**
 * @param {number[][]} matrix
 * @param {number} r
 * @param {number} c
 * @returns {number}
*/

class Solution{

    count_lesser_in_row(arr, mid) {
        let low = 0, high = arr.length - 1
        while(low <= high) {
            let md = parseInt((low + high) / 2)
            if(arr[md] <= mid) {
                low = md + 1
            } else
                high = md - 1
        }
        return low
        
    }


    
    median(matrix,r,c){
        //code here
        //1. binary search to find mid -> return med index
        //2. binary search to find count

        let low = 0, high = Math.pow(10, 9)
        let median_bound = parseInt((r * c) / 2)
        while (low <= high) {
            let mid = parseInt((low + high) / 2) //?
            let cnt = 0
            for (let i = 0; i <= r - 1; i++) {
                cnt = cnt + this.count_lesser_in_row(matrix[i], mid)
            }
            //console.log(cnt)
            if (cnt <= median_bound) {
                low = mid + 1
            } else {
                high = mid - 1
            }

            
            
        }

        return low
        
    }
}


let solution = new Solution()
let R = 3, C = 1
let mat = [[1], [2], [3]]

console.log(solution.median(mat, R, C))
