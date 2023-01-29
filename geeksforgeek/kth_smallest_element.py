class Solution:
    def kthSmallest(self,arr, l, r, k):
        '''
        arr : given array
        l : starting index of the array i.e 0
        r : ending index of the array i.e size-1
        k : find kth smallest element and return using this function
        '''
        arr.sort()
        return arr[k-1]

if __name__ == '__main__':
    arr =  [7, 10, 4, 3, 20, 15]
    n = len(arr)
    k = 3
    s = Solution()
    print(s.findMidSum(arr, 0, n-1, k))
