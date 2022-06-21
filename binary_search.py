

class Solution:

    def _search(self, arr, l, r, k):
        if r >= l:
            #print("l: %d, r: %d" %(l,r))
            mid = (l + r) // 2
            #print(mid)
            if arr[mid] == k:
                return mid
            elif arr[mid] > k:
                return self._search(arr, l, mid - 1, k)
            else:
                return self._search(arr, mid + 1, r, k) 
        else:
            return -1
 

    def binarysearch(self, arr, n, k):
        return self._search(arr, 0, n - 1, k)
 
if __name__ == '__main__':
    arr  = [1, 2, 3, 4, 5] 
    n = len(arr)
    k = 6
    s = Solution()
    print(s.binarysearch(arr, n, k))
