def isSubset( a1, a2, n, m):
    hm = set()
    for _a in a1:
        hm.add(_a)
    for _a in a2:
        if _a not in hm:
            return "No"
    return "Yes"


if __name__ == '__main__':
    arr1 = [11, 1, 13, 21, 3, 7]
    arr2 = [11, 3, 7, 1]
    n,m = len(arr1), len(arr2)
    print(isSubset( a1, a2, n, m))
        
