class Solution:
    def merge(self,arr, l, m, r): 
        # code here
        #print("BEGIN")
        #print(arr)
        #print("%d, %d, %d" %(l,m,r))
        L = arr[l:m+1]
        R = arr[m+1:r+1]
        if len(L) < 1 or len(R) < 1:
            return
        #print(L)
        #print(R)
        i,j= 0,0
        k = l
        while (i < len(L)) and (j<len(R)):
            if L[i] < R[j]:
                arr[k] = L[i]
                i+=1
            else:
                arr[k] = R[j]
                j+=1
            k+=1


        #print(arr)

        while j < len(R):
            arr[k] = R[j]
            k+=1
            j+=1

        

        while i < len(L):
            arr[k] = L[i]
            k+=1
            i+=1

        
            
            

    def mergeSort(self,arr, l, r):
        #code here
        if l < r:
            mid = (l + r) // 2
            self.mergeSort(arr, l, mid)
            self.mergeSort(arr, mid+1, r)
            self.merge(arr, l, mid, r)
            
        return arr

if __name__ == '__main__':
    arr  = [10, 9, 8, 7, 6, 5, 4, 3, 2, 1] 
    n = len(arr)
    s = Solution()
    print(s.mergeSort(arr, 0, n-1))
