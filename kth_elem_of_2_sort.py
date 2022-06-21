class Solution:

    def kthElement(self, arr1, arr2, n, m, k):

        #base case

        if n > m:
            self.kthElement(arr2, arr1, m, n, k)

        #print(k)
        
        if n == 0:
            #print(arr2[k-1])
            return arr2[k-1]

        #if (k == 1):
        #    return min(arr1[0], arr2[0])
        
        #general case

        mid1 = (n - 1) // 2
        mid2 = (m -1) // 2
        print("mid1: %d, mid2: %d" %(mid1, mid2))
        if arr1[mid1] < arr2[mid2]:
           return self.kthElement(arr1[mid1:],arr2,  n - 1 - mid1, m, k - mid1)
        else:
           return self.kthElement(arr1,  arr2[mid2:], n, m - 1 - mid2, k  - mid2)

    
if __name__ == '__main__':
    arr1 = [2, 3, 6, 7, 9]
    arr2 = [1, 4, 8, 10]
    k = 5
    n = len(arr1)
    m = len(arr2)
    s = Solution()
    print(s.kthElement(arr1, arr2, n, m, k))
