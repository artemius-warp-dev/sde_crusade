#User function Template for python3

class Solution:
    def f(self, i, target, coins, dp):
        if target == 0:
            dp[i][target] = 1
            return dp[i][target]
        if i ==0:
            return 0
        if(dp[i][target]) != -1:
            return dp[i][target]
        
        
        no_take = self.f(i-1, target, coins, dp)
        take = 0
        
        if(coins[i-1] <= target):
            take = self.f(i, target - coins[i-1], coins, dp)
        
        dp[i][target] = take + no_take
        return dp[i][target]
        
    
    def count(self, coins, N, Sum):
        # code here
        dp = [[-1 for i in range(Sum+1)] for j in range(N+1)]
        #print(dp)
        r = self.f(N, Sum, coins, dp)
        return r
        
        
