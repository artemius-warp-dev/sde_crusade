class Solution:
    #User function Template for python3
    
    # arr[]: Input Array
    # N : Size of the Array arr[]
    #Function to count inversions in the array.
    def inversionCount(self, arr, n):
        # Your Code Here
        return self.mergeSort(arr, 0, n-1)


    def merge(self,arr, l, m, r): 
        # code here
        icount = 0
        L = arr[l:m+1]
        R = arr[m+1:r+1]
        l_size, r_size = len(L), len(R)
        i,j= 0,0
        k = l
        while (i < l_size) and (j<r_size):
            if L[i] <= R[j]:
                arr[k] = L[i]
                i+=1
            else:
                icount+= l_size - i
                arr[k] = R[j]
                j+=1
            k+=1

        while j < r_size:
            arr[k] = R[j]
            k+=1
            j+=1

        while i < l_size:
            arr[k] = L[i]
            k+=1
            i+=1
        return icount
            

    def mergeSort(self,arr, l, r):
        if l < r:
            mid = (l + r) // 2
            return self.mergeSort(arr, l, mid) + self.mergeSort(arr, mid+1, r) +  self.merge(arr, l, mid, r)
        return 0

        

if __name__ == '__main__':
    arr =  [468,335,1,170,225,479,359,463,465,206,146,282,328,462,492,496,443,328,437,392,105,403,154,293,383,422,217,219,396,448,227,272,39,370,413,168,300,36,395,204,312,323]
    n = len(arr)
    s = Solution()
    print(s.inversionCount(arr, n))
