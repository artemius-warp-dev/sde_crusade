class Solution:
    def relativeSort (self,A1, N, A2, M):
        hm = {}
        res = []
        
        for e in A1:
            if e not in hm:
                hm[e] = 1
            else:
                hm[e]+=1
        for e in A2:
            if e not in hm:
                continue
            res.extend([e] * hm[e])

            hm[e] = 0

        rem = []
        for key in hm:
            if hm[key] != 0:
                rem.append(key)

        rem.sort()

        for e in rem:
            res.extend([e] * hm[e])

        return res

        
        


if __name__ == '__main__':
    A1 = [15, 47, 59, 326, 124, 5, 623, 128, 483, 153, 125, 232, 154, 555, 656, 485, 659, 32, 125, 326, 324, 96, 565, 154, 112, 32, 85, 563, 32, 481, 32]
    A2 = [20, 47, 20, 5, 125, 154, 555, 32, 324]
    N = len(A1)
    M = len(A2)
    s = Solution()
    print(s.relativeSort(A1, N, A2, M))
