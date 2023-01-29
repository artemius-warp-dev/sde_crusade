from collections import defaultdict

class Solution:
    def check(self, A, B, N):
        count = defaultdict(int)

        for elem in A:
            count[elem] += 1

        for elem in B:
            if count[elem] == 0:
                return 0
            else:
                count[elem] -= 1
        return 1
        
if __name__ == '__main__':
    A = [1,2,5,4,0]
    B = [2,4,5,0,1]
    size = len(A)
    s = Solution()
    print(s.check(A, B, size))
