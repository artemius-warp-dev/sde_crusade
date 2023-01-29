class Solution
{
    //Function to search a given number in row-column sorted matrix.
    search(matrix, n, m, x)
    {
    	// code here
        let irow = 0, icolumn = m - 1
        while (irow < n && icolumn > 0) {
            if (matrix[irow][icolumn] == x) {
                return true
            } 
            if (matrix[irow][icolumn] > x) {
                icolumn--
            } else {
                irow++
            }
        }
        return false
    }
}

let solution = new Solution()
let mat =    [[18, 21, 27, 38, 55, 67]]
console.log(solution.search(mat, 1, 6, 55))
