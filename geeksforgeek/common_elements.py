from collections import defaultdict

class Solution:
    def common_element(self,v1,v2):
        hm = defaultdict(int)
        common_elements = []
        for v1_item in v1:
            hm[v1_item] += 1

        for v2_item in v2:
            if hm[v2_item] > 0:
                common_elements.append(v2_item)
                hm[v2_item] -= 1

        common_elements.sort()
        return common_elements

        



if __name__ == '__main__':
    arr1 = [3, 4, 2, 2, 4]
    arr2 = [3, 2, 2, 7]
    n1 = len(arr1)
    n2 = len(arr2)
    s = Solution()
    print(s.common_element(arr1, arr2))
