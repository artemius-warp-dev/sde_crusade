class Solution
{
    first(arr, low, high) {
        if (high >= low) {
            let mid = low + (high - low)/2
            //base case
            if ((arr[mid] == 1 && (arr[mid-1] == 0 || mid == 0))) {
                return mid
            } else if (arr[mid] == 0) {
                return this.first(arr, mid + 1, high)
            } else {
                return this.first(arr, low, high - 1)
            }
        }


        
        return -1;
    }

    maxOnes(mat, n, m){
       //code here
        let max_count = 0, max_row_index = 0
        let last_column_index = m -1
        let first_occurence, count_in_row

        for(let i = 0; i < n; i++) {
            
            first_occurence = this.first(mat[i], 0, last_column_index)
            if (first_occurence != -1) {
                count_in_row = last_column_index - first_occurence
                //console.log(count_in_row)
                if (count_in_row > max_count) {
                    max_count = count_in_row
                    max_row_index = i
                }
            }
            
        }
        return max_row_index

       
    }
}

let solution = new Solution()
let N = 10, M = 3
let mat =    [[0, 0, 1], [0, 1, 1], [0, 1, 1], [1, 1, 1], [0, 0, 1], [0, 0, 1], [0, 1 ,1], [0, 0, 1], [0, 0, 0], [1, 1, 1]]

console.log(solution.maxOnes(mat, N, M))
