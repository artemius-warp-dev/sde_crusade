import bisect
class Solution:
    

    def binary_search(self,arr, l,r,x, ans):
        if r >=l:
            mid = (l+r) // 2
            mid_val = arr[mid]
            
            if mid_val > x:
                ans = mid
                return self.binary_search(arr, l, mid - 1, x, ans)
            else:
                return self.binary_search(arr, mid+1, r, x, ans)

        return ans
    
    def _count(self,x, arr, N, NoOfY):
        if x == 0:
            return 0

        if x == 1:
            return NoOfY[0]

        #idx = self.binary_search(arr, 0, N-1, x, -1) can't go through time limit
        idx = bisect.bisect_right(arr, x)
        if idx == -1:
            ans = 0
        else:
            ans = N - idx

        ans += NoOfY[0] + NoOfY[1]

        if x == 2:
            ans -= NoOfY[3] + NoOfY[4]
 
        if x == 3:
            ans += NoOfY[2]
 
        return ans
    
     #Function to count number of pairs such that x^y is greater than y^x.     
    def countPairs(self,a,b,M,N):
        #code here
        b.sort()
        NoOfB = [0] * 5
        for i in range(N):
            if b[i] < 5:
                NoOfB[b[i]] += 1
        total_pairs = 0

        for i in range(M):
           total_pairs += self._count(a[i], b, N, NoOfB)

        return total_pairs
                


if __name__ == '__main__':
    arr1  = [2, 3, 4, 5]
    arr2 =  [1, 2, 3]
    m = len(arr1)
    n = len(arr2)
    s = Solution()
    print(s.countPairs(arr1,arr2,m,n))
