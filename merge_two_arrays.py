class Solution:

    #Function to merge the arrays.
    def merge(self,arr1,arr2,n,m):
        #code here
        i,j = n -1, 0
        
        while i>-1 and j < m:
            if arr1[i] > arr2[j]:
                arr1[i], arr2[j] = arr2[j], arr1[i]
            i-=1
            j+=1
                
        arr1.sort()
        arr2.sort()
    
    # #Function to merge the arrays.
    # def merge(self,arr1,arr2,n,m):
    #     #code here
    #     for i in range(n):
    #         if arr1[i] > arr2[0]:
    #             arr1[i], arr2[0] = arr2[0], arr1[i]
    #             arr2.sort()
            

    #     return arr1, arr2


if __name__ == '__main__':
    arr1  = [1, 3, 5, 7] 
    arr2  = [0, 2, 6, 8, 9]
    
    n = len(arr1)
    m = len(arr2)

    s = Solution()
    print(s.merge(arr1, arr2, n, m))
