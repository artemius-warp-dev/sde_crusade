// User function Template for javascript

/**
 * @param {number} V
 * @param {number[][]} adj
 * @returns {number[]}
*/
class Solution {
    // Function to return a list containing the DFS traversal of the graph.
    dfsOfGraph(V, adj) {
        // code here
        let list = new Array()
        let visited = new Array(V).fill(false)
        this.dfs(adj, 0, list, visited)
        return list
    }
    
    dfs(adj, s, list, visited) {
        visited[s] = true
        list.push(s)
        for(const u of adj[s]) {
            //console.log(u)
            if(visited[u] == false) {
                this.dfs(adj, u, list, visited)
            }
        }
    }
}



