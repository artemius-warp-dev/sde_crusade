from collections import defaultdict
class Solution:
    def countDistinct(self, A, N, K):
        # Code here
        hm = defaultdict(int)
        res = []
        distinct_numbers = 0
        for i in range(K):
            if hm[A[i]] == 0:
                distinct_numbers+=1
            hm[A[i]] +=1

        res.append(distinct_numbers)
            
        for i in range(K, N):
            if hm[A[i-k]] == 1:
                distinct_numbers -=1
            hm[A[i-k]] -=1

            if hm[A[i]] == 0:
                distinct_numbers+=1

            hm[A[i]] +=1

            res.append(distinct_numbers)

            

        return res


if __name__ == '__main__':
    arr  = [4,1,1] 
    n = len(arr)
    k = 2
    s = Solution()
    print(s.countDistinct(arr, n, k))
        
