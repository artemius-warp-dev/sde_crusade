
class Solution:
    #Back-end complete function Template for Python 3
    
    #Function to find the leaders in the array.
    def leaders(self, A, N):
        #Code here
        output_arr = []
        max_right_elem = A[N-1]
        output_arr.append(max_right_elem)
        for i in range(N-2, -1, -1):
            if A[i] >= max_right_elem:
                max_right_elem = A[i]
                output_arr.append(max_right_elem)
        output_arr.reverse()
        return output_arr

#{ 
#  Driver Code Starts
#Initial Template for Python 3

import math


    
def main():
    
    T=int(input())
    
    while(T>0):
        
        
        N=int(input())
        
        A=[int(x) for x in input().strip().split()]
        obj = Solution()
        
        A=obj.leaders(A,N)
        
        for i in A:
            print(i,end=" ")
        print()
        
        T-=1

if __name__=="__main__":
    main()
# } Driver Code Ends
