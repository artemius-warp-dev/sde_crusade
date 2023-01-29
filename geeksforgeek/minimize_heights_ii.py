class Solution:
    def getMinDiff(self, arr, n, k):
        # code here
        arr.sort()
        a = arr[0]
        b = arr[n-1]
        diff = b - a
        for i in range(1, n):
            y = arr[i]
            x = arr[i-1]
            if(y-k >= 0):
                min_value = min(a+k, y - k)
                max_value = max(b-k, x + k)
                diff = min(diff, max_value - min_value)

        return diff
