class Solution:

    def get_sum(self, a):
        _sum = 0
        for e in a:
            _sum += e
        return _sum

    def get_target(self, a,b):
        sum1 = self.get_sum(a)
        sum2 = self.get_sum(b)
        return (sum1 - sum2) / 2

    def findSwapValues(self,a, n, b, m):
        a.sort()
        b.sort()

        target = self.get_target(a,b)
        i,j = 0,0
        while i < n and j<m:
            diff = a[i] - b[j]
            if diff == target:
                return 1
            elif diff < target:
                i+=1
            else:
                j+=1
        return -1
                

if __name__ == '__main__':
    arr =  [10, 10, 10, 10]
    arr2 = [5, 5, 5, 5]
    n = len(arr)
    m = len(arr2)
    s = Solution()
    print(s.findSwapValues(arr, n, arr2, m))
