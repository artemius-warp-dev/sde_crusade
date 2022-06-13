class Solution:
    def minimumPlatform(self,n,arr,dep):
        i,j = 1,0

        arr.sort()
        dep.sort()

        
        platforms = 1
        result = 1
        while i<n and j<n:
            if arr[i] <= dep[j]:
                i+=1
                platforms+=1
            else:
                j+=1
                platforms-=1

            if platforms > result:
                result = platforms
        return result

if __name__ == '__main__':
    arr = [900, 940, 950, 1100, 1500, 1800]
    dep = [910, 1200, 1120, 1130, 1900, 2000]
    n = len(arr)
    s = Solution()
    print(s.minimumPlatform(n, arr, dep))
