class Solution:
    def overlappedInterval(self, Intervals):
	#Code here
        merged_intervals = []
        if len(Intervals) == 0:
            return merged_intervals

        Intervals.sort(key=lambda x: (x[0], x[1]))
       # print(Intervals)
        temp = Intervals[0]
        merged_intervals.append(temp)

        for it in Intervals[1:]:
            if temp[1] >= it[0]:
                temp[1] = max(temp[1],it[1])
            else:
                merged_intervals.append(it)
                temp = merged_intervals[-1]

        return merged_intervals



if __name__ == '__main__':
    intervals = [[6, 8], [1, 9], [2, 4], [4, 7]]
    s = Solution()
    print(s.overlappedInterval(intervals))

