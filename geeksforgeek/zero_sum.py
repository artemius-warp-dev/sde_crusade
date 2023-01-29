class Solution:
    #Function to count subarrays with sum equal to 0.
    def findSubArrays(self,arr,n):
        sum_map = {}
        cur_sum = 0
        out = []
        count = 0
        for i in range(n):

            cur_sum += arr[i]

            if cur_sum == 0:
                out.append((0, i))

            al = []

            if cur_sum in sum_map:
                al = sum_map.get(cur_sum)
            for it in range(len(al)):
                out.append((al[it] + 1, i))
            al.append(i)
            sum_map[cur_sum] = al
                        
           
        count = len(out)
        return count
        
            
        #return: count of sub-arrays having their sum equal to 0

if __name__ == '__main__':
    arr  = [6,-1,-3,4,-2,2,4,6,-12,-7] 
    n = len(arr)
    s = Solution()
    print(s.findSubArrays(arr, n))
