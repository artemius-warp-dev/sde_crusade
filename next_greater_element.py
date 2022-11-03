class Solution:
    def nextLargerElement(self,arr,n):
        #code here
        stack = []
        res = []
        i = n - 1
        while (i >= 0):
            if len(stack) == 0:
                 stack.append(arr[i])
                 res.append(-1)
            elif arr[i] < stack[-1]:
                 res.append(stack[-1])
                 stack.append(arr[i])
            else:
                 while(arr[i] >= stack[-1]):
                     stack.pop()
                     if(len(stack) == 0):
                         stack.append(-1)
                         break
                     
                 
                 res.append(stack[-1])
                 stack.append(arr[i])
                

            i=i-1
        res.reverse()
        return res


if __name__ == '__main__':
    arr = [7, 8, 1, 4]
    #[
    #[8, -1, 4, -1]

    #
    n = len(arr)
    s = Solution()
    print(s.nextLargerElement(arr, n))       
