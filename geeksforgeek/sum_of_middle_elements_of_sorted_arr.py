class Solution:

    def get_mid(self, arr, n):
        if n % 2 == 0:
            return (arr[int(n / 2)] +
                    arr[int(n / 2) - 1]) //2
        else:
            return (arr[int(n/2)] + arr[int(n/2)]) //2

    
    def findMidSum(self, arr1, arr2, n): 
	# code here

        # base cases

        if n == 0:
            return -1
        
        if n == 1:
            return (arr1[0] + arr2[0])

        if n == 2:
            return (max(arr1[0], arr2[0]) + min(arr2[1], arr1[1]))

        # general case

        m1 = self.get_mid(arr1, n)
        m2 = self.get_mid(arr2, n)


        if m1 > m2:
            if n % 2 == 0:
                return self.findMidSum(arr1[:int(n/2) + 1], arr2[int(n/2) - 1:], int(n/2) +1)
            else:
                return self.findMidSum(arr1[:int(n/2)+1], arr2[int(n/2):], int(n/2) + 1)
        else:
            if n % 2 == 0:
                return self.findMidSum(arr1[int(n/2)-1:], arr2[:int(n/2)+1], int(n/2)+1)
            else:
                return self.findMidSum(arr1[int(n/2):], arr2[:int(n/2)+1], int(n/2)+1)
            
        


        
if __name__ == '__main__':
    arr1  = [5,10,12,16,17,18,22,24,27,29,29,29]
    arr2 =  [5, 7, 9, 10, 10, 13, 16, 18, 20, 22, 26, 29]
    n = len(arr1)
    s = Solution()
    print(s.findMidSum(arr1, arr2, n))
