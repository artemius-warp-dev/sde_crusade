class Solution:
    def spirallyTraverse(self, matrix, r, c):
        return self._spirallyTraverse(matrix, 0, 0, r-1, c-1, [])
    def _spirallyTraverse(self, matrix, top, left, bottom, right, output_numbers):
        #base case
        
        if (top > bottom or left > right):
            return output_numbers

        i = top
        while i <= right:
            output_numbers.append(matrix[top][i])
            i+=1
        
        #print last column
        i = top + 1
        while i <= bottom:
            output_numbers.append(matrix[i][right])
            i+=1
        
        # #print bottom row
        if(bottom != top):
            i = right - 1
            while i >= left:
                output_numbers.append(matrix[bottom][i])
                i-=1
        
        #print first column
        if(right != left):
            i = bottom - 1
            while i >= top+1:
                output_numbers.append(matrix[i][left])
                i-=1
        return self._spirallyTraverse(matrix, top+1, left+1, bottom-1, right-1, output_numbers)
 

if __name__ == '__main__':
    #matrix  = [[1, 2, 3, 4],
    #           [5, 6, 7, 8],
    #           [9, 10, 11, 12],
    #           [13, 14, 15, 16]]

    matrix = [
        [1, 2, 3, 4],
        [5, 6, 7, 8],
        [9, 10, 11, 12]
    ]
    
    #matrix = [[1,2], [3,4], [5,6]]
    # 1,2
    # 3,4
    # 5,6

    r = len(matrix)
    c = len(matrix[0])
    s = Solution()
    print(s.spirallyTraverse(matrix, r, c))
