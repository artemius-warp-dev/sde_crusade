class Solution:
    def fourSum(self, arr, k):
        # Sort the array in increasing order,
        # using library function for quick sort
        A = arr
        n = len(A)
        X = k
        A.sort()
        result = set()
        # Now fix the first 2 elements one by
        # one and find the other two elements
        for i in range(n - 3):
            for j in range(i + 1, n - 2):

                # Initialize two variables as indexes
                # of the first and last elements in
                # the remaining elements
                l = j + 1
                r = n - 1

                # To find the remaining two elements,
                # move the index variables (l & r)
                # toward each other.
                while (l < r):
                    s = A[i] + A[j] + A[l] + A[r]
                    if(s == X):

                        result.add((A[i], A[j], A[l], A[r]))
                        
                        l += 1
                        r -= 1

                    elif (s < X):
                        l += 1
                    else: # A[i] + A[j] + A[l] + A[r] > X
                        r -= 1
        res = list(result)
        res.sort()
        return res
    # def fourSum(self, arr, k):
    #     arr.sort()
    #     size = len(arr)
    #     pair_sum = {}
    #     response = {}
    #     for i in range(size - 1):
    #         for j in range(i+1, size):
    #             pair_sum[arr[i] + arr[j]] = [i,j]


    #     for i in range(size-1):
    #         for j in range(i+1, size):
    #             _sum = arr[i] + arr[j]
    #             k_sum = k - _sum
    #             if k_sum in pair_sum:
    #                 pair_sum_values =  pair_sum[k_sum]
    #                 x,y = pair_sum_values[0], pair_sum_values[1]
    #                 if x != i and x != j and y != i and y != j:
    #                    variant = [arr[i] , arr[j], arr[x],arr[y]]
    #                    variant.sort()
    #                    str_key = str(variant)

    #                    if str_key not in response:
    #                        response[str_key] = variant
                       
    #     res = []
    #     for variant in response.values():
    #         res.append(variant)

    #     res.sort()
    #     return res

if __name__ == '__main__':
    arr1 = [10,2,3,4,5,7,8]
    K = 23
    s = Solution()
    print(s.fourSum(arr1, K))
