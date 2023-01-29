def sortByFreq(arr, n):
    hm = {}

    for e in arr:
        if e in hm:
            hm[e] +=1
        else:
            hm[e] = 1

    arr.sort(key=lambda x: (-hm[x], x))
    return arr
    #print(hm)




if __name__ == "__main__":
    arr = [2, 5, 2, 6, -1, 9999999, 5, 8, 8, 8]
    n = len(arr)
    solution = sortByFreq(arr, n)
    print(*solution)

#THIS IS ACCEPTED SOLUTION FOR GFG
# #code
# def sortByFreq(arr, n):
#     hm = {}

#     for e in arr:
#         if e in hm:
#             hm[e] +=1
#         else:
#             hm[e] = 1

#     arr.sort(key=lambda x: (-hm[x], x))
#     return arr

# k = int(input())
# for i in range(k):
#     n = int(input())
#     arr = list(map(int, input().strip().split()))[:n]
#     arr = sortByFreq(arr, n)
#     print(*arr)

