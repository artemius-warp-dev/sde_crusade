class Solution:	
    #Function to reverse every sub-array group of size k.
	def reverseInGroups(self, arr, N, K):
		# code here
		i = 0
		while(i<N):
		    left = i
		    right = min(i+K-1, N-1)
		    
		    while left < right:
		        arr[left], arr[right] = arr[right], arr[left]
		        left+=1
		        right-=1
		        
		    i+=K
		return arr



if __name__ == '__main__':
    arr = [0,2,1,2,0]
    size = len(arr)
    s = Solution()
    print(s.reverseInGroups(arr, size, 3))
