class Solution:
    def firstElementKTime(self,  a, n, k):
        # code here
        counter_map = {}

        for i in range(n):
            if a[i] in counter_map:
                counter_map[a[i]] += 1
            else:
                counter_map[a[i]] = 1

            if counter_map[a[i]] == k:
                return a[i]

        return -1
  
                

    

if __name__ == '__main__':
    arr =  [3, 4, 1, 3, 4, 4]
    n = len(arr)
    k = 2
    s = Solution()
    print(s.firstElementKTime(arr, n,k))
