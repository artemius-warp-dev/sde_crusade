class Solution:
    def allPairs(self, A, B, N, M, X):
        # Your code goes here
        
        hm = set()
        for a_item in A:
            hm.add(a_item)

        output = []
        for b_item in B:
            _a_item = X - b_item
            if _a_item in hm:
                output.append([_a_item, b_item])
        output.sort()
        return output


if __name__ == '__main__':
    arr1 =  [1, 2, 4, 5, 7]
    arr2 = [5, 6, 3, 4, 8]
    N = len(arr1)
    M = len(arr2)
    X = 9
    s = Solution()
    s.allPairs(arr1, arr2, N, M, X)
