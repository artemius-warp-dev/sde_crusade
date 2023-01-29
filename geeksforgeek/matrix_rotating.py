class Solution:
    def rotateMatrix(self, M, N, Mat):
        #indexes
        mat = Mat
        top = 0
        bottom = len(mat) - 1
        left = 0
        right = len(mat[0]) - 1
        while left < right and top < bottom:
            prev = mat[top+1][left]

            for i in range(left, right+1):
                curr = mat[top][i]
                mat[top][i] = prev
                prev = curr

            top += 1

            for i in range(top, bottom+1):
                curr = mat[i][right]
                mat[i][right] = prev
                prev = curr

            right -= 1

            for i in range(right, left-1, -1):
                curr = mat[bottom][i]
                mat[bottom][i] = prev
                prev = curr
            bottom -= 1

            for i in range(bottom, top-1, -1):
                curr = mat[i][left]
                mat[i][left] = prev
                prev = curr
            left +=1

        return mat
            

    # Utility Function
    def printMatrix(self, mat):
        for row in mat:
            print(row)               

    
if __name__ == '__main__':
    M= 3
    N= 3
    Mat=[[1,2,3],[4,5,6],[7,8,9]]
    s = Solution()
    matrix = s.rotateMatrix(M, N, Mat)
    print()
    s.printMatrix(matrix)
        
