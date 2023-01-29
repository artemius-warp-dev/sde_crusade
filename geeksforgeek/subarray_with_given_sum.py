class Solution:
    
    def subArraySum(self,arr,n,s):
      current_sum = arr[0]
      slide_window_p = 0

      i = 1
      while i <= n:
          #print("current_sum %d and position: %d" %(current_sum, i))
          while current_sum > s:
              #print("array elem from slide pointer: %d" %(arr[slide_window_p]))
              current_sum = current_sum - arr[slide_window_p]
              slide_window_p+=1

          if current_sum == s:
              return [slide_window_p+1, i]

          if i < n:
              current_sum += arr[i]
          i+=1
      return [-1]

if __name__ == '__main__':
    arr = [135, 101, 170, 125, 79, 159, 163, 65, 106, 146, 82, 28, 162, 92, 196, 143, 28, 37, 192, 5, 103, 154, 93, 183, 22, 117, 119, 96, 48, 127, 172, 139, 70, 113, 68, 100, 36, 95, 104, 12 ,123, 134]
    #arr = [1,2,3,7,5]
    size = len(arr)
    sum_ = 468
    s = Solution()
    print(s.subArraySum(arr, size, sum_))
