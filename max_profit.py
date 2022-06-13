class Solution:
    def stockBuySell(self, A, n):
        i = 0
        result = []
        while i < n - 1:

            while (i < n-1) and (A[i+1] <= A[i]):
                i+=1


            if i == n - 1:
                break

            buy = i
            i+=1
            while (i < n) and (A[i] >= A[i-1]):
                i+=1
                
            sell = i - 1
            
            result += [[buy, sell]]
        return result

            

if __name__ == '__main__':
    price = [11, 42, 49, 96, 23, 20, 49, 26, 26, 18, 73, 2, 53, 59, 34, 99, 25, 2]
    n = len(price)
    s = Solution()
    print(s.stockBuySell(price, n))
