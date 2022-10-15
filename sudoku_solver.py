class Solution:
    #Function to find a solved Sudoku. 
    def SolveSudoku(self,grid):
        for i in range(len(grid)):
            for j in range(len(grid)):
                if grid[i][j] == 0:
                    for val in range(1,10):
                        if self.isValid(i,j,val,grid):
                            grid[i][j] = val
                        
                            if self.SolveSudoku(grid) is True:
                                return True
                                
                            else:
                                grid[i][j] = 0
                    return False
        return True            
        
    
    def isValid(self,row,col,val,grid):
        for i in range(len(grid)):
            if grid[i][col] == val:
                return False
    
            if grid[row][i] ==val:
                return False
    
            if grid[3 * (row // 3) + i // 3][3 * (col // 3) + i % 3] == val:
                return False
        return True
        
    #Function to print grids of the Sudoku.    
    def printGrid(self,arr):
        for i in arr:
            for j in i:
                print(j,end=' ')
