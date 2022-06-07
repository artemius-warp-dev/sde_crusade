from sys import maxsize

class Solution:
    
    def maxSubArraySum(self,arr,N):
        max_sum = arr[0]
        current_sum = 0

        for i in range(0, N):
            current_sum+= arr[i]
           

            if current_sum > max_sum:
                max_sum = current_sum
            if current_sum < 0:
                current_sum = 0

        return max_sum

if __name__ == '__main__':
    arr = [-47, 43, 94, -94, -93, -59, 31, -86]
    #arr = [1,2,3,-2,5]
    #arr = [-1,-2,-3,-4]
    size = len(arr)
    s = Solution()
    print(s.maxSubArraySum(arr, size))
