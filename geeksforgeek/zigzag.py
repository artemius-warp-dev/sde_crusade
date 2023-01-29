class Solution:	
    #Function to reverse every sub-array group of size k.
     def zigZag(self, arr, n):
         flag = True
         for i in range(n-1):
             if flag is True: # <
                 if arr[i] > arr[i+1]:
                     arr[i], arr[i+1] = arr[i+1], arr[i]
             else: # >
                 if arr[i] < arr[i+1]:
                     arr[i], arr[i+1] = arr[i+1], arr[i]
             flag = bool(1 - flag)
         return arr



if __name__ == '__main__':
    arr = [4,3,7,8,6,2,1]
    size = len(arr)
    s = Solution()
    print(s.zigZag(arr, size))
