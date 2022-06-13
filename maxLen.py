class Solution:

    def maxLen(self, n, arr):
        hm = {}
        cur_sum = 0
        max_len = 0
        for i in range(n):
            cur_sum += arr[i]

            if cur_sum == 0:
                max_len = i + 1
            
            if cur_sum in hm:
                max_len = max(max_len, i - hm[cur_sum])
            else:
                hm[cur_sum] = i
        return max_len

if __name__ == '__main__':
    arr =  [-1, 1, -1, 1]
    n = len(arr)
    s = Solution()
    print(s.maxLen(n, arr))
