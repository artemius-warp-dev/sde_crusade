class Solution:
    

    def sort012(self, arr, n):
        count0, count1, count2 = 0,0,0

        for i in range(n):
            if arr[i] == 0:
                count0 += 1
            elif arr[i] == 1:
                count1 += 1
            elif arr[i] == 2:
                count2 += 1
        i = 0

        while count0 > 0:
            arr[i] = 0
            i+=1
            count0 -= 1
        while count1 > 0:
            arr[i] = 1
            i+=1
            count1-=1
        while count2 > 0:
            arr[i] = 2
            i+=1
            count2-=1

        return arr
    

    # def sort012(self, arr, n):
    #     lo,hi,mid = 0, n-1, 0
    #     while mid <= hi:
    #         if(arr[mid] == 0):
    #             arr[lo], arr[mid] = arr[mid], arr[lo]
    #             lo+= 1
    #             mid+= 1
    #         elif(arr[mid] == 1):
    #             mid+= 1
    #         else:
    #             arr[mid], arr[hi] = arr[hi], arr[mid]
    #             mid+= 1
    #             hi-= 1
    #     return arr

if __name__ == '__main__':
    arr = [0,2,1,2,0]
    size = len(arr)
    s = Solution()
    print(s.sort012(arr, size))
