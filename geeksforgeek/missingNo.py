class Solution:
    def missingNumber(self, array,n):
        x1 = array[0]; x2 = 0

        for i in range(1, n-1):
            x1 = x1 ^ array[i]

        for i in range(1, n+1):
            
            x2 = x2 ^ i
        return x1 ^ x2

if __name__ == '__main__':
    arr = [1,2,3,5]
    N = 5
    s = Solution()
    print(s.missingNumber(arr, N))
