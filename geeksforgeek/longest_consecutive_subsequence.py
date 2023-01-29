class Solution:
    
    # arr[] : the input array
    # N : size of the array arr[]
    
    #Function to return length of longest subsequence of consecutive integers.
    def findLongestConseqSubseq(self,arr, N):
        #code here
        hash_map = set()
        count = 0
        for elem in arr:
            hash_map.add(elem)

        for i in range(N):
            if arr[i] - 1 not in hash_map:
                start_point = arr[i]

                while start_point in hash_map:
                    start_point+=1
                count = max(count, start_point - arr[i])
        return count

if __name__ == '__main__':
    arr  = [1,9,3,10,4,20,2] 
    n = len(arr)
    s = Solution()
    print(s.findLongestConseqSubseq(arr, n))
