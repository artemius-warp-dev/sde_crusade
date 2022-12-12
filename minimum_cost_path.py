from heapq import heapify,heappush,heappop
class Solution:
	def minimumCostPath(self, grid):
	    n = len(grid)
	    heap = [(grid[0][0],0,0)]
	    heapify(heap)
	    v = set((0,0))
	    while heap:
	        val,i,j = heappop(heap)
	        if i==n-1 and j==n-1:
	            return val
	           
	        if i<n-1 and (i+1,j) not in v:
	            heappush(heap,(val+grid[i+1][j],i+1,j))
	            v.add((i+1,j))
	        if i>0 and (i-1,j) not in v:
	            heappush(heap,(val+grid[i-1][j],i-1,j))
	            v.add((i-1,j))
	        if j<n-1 and (i,j+1) not in v:
	            heappush(heap,(val+grid[i][j+1],i,j+1))
	            v.add((i,j+1))
	        if j>0 and (i,j-1) not in v:
	            heappush(heap,(val+grid[i][j-1],i,j-1))
	            v.add((i,j-1))
	       
	    return val
