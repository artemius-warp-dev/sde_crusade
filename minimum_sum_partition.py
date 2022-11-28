#User function Template for python3
import sys
class Solution:
	def lcs():
	    pass
	
	def minDifference(self, arr, n):
		# code here
		total_sum = 0
		for e in arr:
		    total_sum = total_sum + e
		
		rows = n
        cols = total_sum  + 1
        
        dp = [[False for _ in range(cols)] for _ in range(rows)]
		
		for i in range(n):
		    dp[i][0] = True
		if(arr[0] <= total_sum):
		    dp[0][arr[0]] = True
		    
		for i in range(1,n):
		    for target in range(1, total_sum):
		        no_take = dp[i-1][target]
		        take = False
		        if(arr[i] <= target):
		            take = dp[i-1][target - arr[i]]
		        
		        dp[i][target] = (take or no_take)
		        
		mini = sys.maxsize
		for i in range(0, total_sum//2+1):
		    if(dp[n-1][i] == True):
		        mini = min(mini, abs((total_sum - i) - i))
		
		return mini

