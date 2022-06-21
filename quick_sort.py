


class Solution:

    def quickSort(self,arr,low,high):
        if low < high:
            pi = self.partition(arr, low, high)
            #print(pi)
            self.quickSort(arr, low, pi-1)
            self.quickSort(arr, pi+1, high)

        return arr
        


    def partition(self,arr,low,high):
        #print("%d, %d" %(low, high))
        pivot = arr[high]
        i = low - 1
        #print(high)
        for j in range(low, high):
            if arr[j] <= pivot:
                i+=1
                arr[i], arr[j] = arr[j], arr[i]

        arr[high], arr[i+1] = arr[i+1], arr[high]
        #print(arr)
        return i + 1
 
if __name__ == '__main__':
    arr  = [3,4,5,1,3,0] 
    n = len(arr)
    s = Solution()
    print(s.quickSort(arr, 0, n-1))
