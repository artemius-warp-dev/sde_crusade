#User function Template for python3


class Solution:
    
    
    def isCyclicUtil(self, v, visited, recStack, adj):
        # Mark current node as visited and
        # adds to recursion stack
        visited[v] = True
        recStack[v] = True
        
        # Recur for all neighbours
        # if any neighbour is visited and in
        # recStack then graph is cyclic
        for neighbour in adj[v]:
            if visited[neighbour] == False:
                if self.isCyclicUtil(neighbour, visited, recStack, adj) == True:
                    return True
            elif recStack[neighbour] == True:
                return True
        
        # The node needs to be popped from
        # recursion stack before function ends
        recStack[v] = False
        return False
    
    
    #Function to detect cycle in a directed graph.
    def isCyclic(self, V, adj):
        # code here
        visited = [False] * (V + 1)
        recStack = [False] * (V + 1)
        for node in range(V):
            if visited[node] == False:
                if self.isCyclicUtil(node,visited,recStack, adj) == True:
                    return True
        return False
