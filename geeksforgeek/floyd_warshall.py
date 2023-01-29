class Solution:
	def shortest_distance(self, matrix):
		#Code here
		n = len(matrix)
		
		
		
		for via in range(n):
		    for i in range(n):
		        for j in range(n):
		            if(matrix[i][via] == -1 or matrix[via][j] == -1):
		                continue
		            if(matrix[i][j] == -1 or matrix[i][j] > matrix[i][via] + matrix[via][j]):
		                 matrix[i][j] =  matrix[i][via] + matrix[via][j]
		
