class Solution:	
    #Function to reverse every sub-array group of size k.
     def rearrange(self, arr, n):
         temp = n*[None]

         small, large = 0, n-1

         flag = True
         for i in range(n):
            if flag is True: # large
                 temp[i] = arr[large]
                 large-=1
            else:
                temp[i] = arr[small]
                small+=1
            flag = bool(1-flag)

         for i in range(n):
            arr[i] = temp[i]
         return arr


if __name__ == '__main__':
    arr = [1,2,3,4,5,6,7,8]
    size = len(arr)
    s = Solution()
    print(s.rearrange(arr, size))
