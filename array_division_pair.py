from collections import defaultdict
class Solution:
    def canPair(self, nuns, k):
	# Code here
        hm = defaultdict(int)
        for element in nuns:
            hm[element % k] +=1


        if hm[0] % 2 != 0:
           return False


        #print(hm)
        
        for i in range(1, k // 2 + 1):
            #(9,3), (7, 5)
            if hm[i] != hm[k-i]:
                return False
        return True
                

    
if __name__ == '__main__':
    arr  = [9, 5, 7, 3] 
    k = 6
    s = Solution()
    print(s.canPair(arr, k))
        
