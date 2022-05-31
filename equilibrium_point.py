class Solution:
    
    def equilibriumPoint(self, A,N):

        
        i,start,end, left_sum, right_sum = 0,0,N-1, 0,0
        for i in range(N):
            if (start == end and left_sum == right_sum):
                return start + 1

            if (start == end):
                return -1

            if(left_sum > right_sum):
                right_sum += A[end]
                end-=1
            elif(right_sum > left_sum):
                left_sum += A[start]
                start+=1
            else:
                right_sum += A[end]
                end-=1
        

if __name__ == '__main__':
    arr = [1,3,5,2,2]
    size = len(arr)
    s = Solution()
    print(s.equilibriumPoint(arr, size))
