class Solution:
    def findOnce(self,arr : list, n : int):
        # Complete this function


        def search(arr, l, r):
            #base case
            if l > r:
                return None
            if l == r:
                return arr[l]

            mid = (l+r) // 2

            if mid % 2 == 0:
                if arr[mid] == arr[mid+1]:
                    return search(arr, mid+2, r)
                else:
                    return search(arr, l, mid)
            else:
                if arr[mid] == arr[mid-1]:
                    return search(arr, mid+1, r)
                else:
                    return search(arr, l, mid)
                    
            
        return search(arr, 0, n-1)
    

if __name__ == '__main__':
    arr =  [1, 1, 2, 2, 3, 3, 4, 50, 50, 65, 65]
    n = len(arr)
    s = Solution()
    print(s.findOnce(arr, n))
        
