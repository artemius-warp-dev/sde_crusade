// User function Template for javascript

/**
 * @param {number} V
 * @param {number[][]} adj
 * @returns {boolean}
*/
class Solution {
    // Function to detect cycle in an undirected graph.
    isCycle(V, adj) {
        // code here
        let vis = new Array(V+1).fill(false)
        
        
        let is_cycle = function(n_index, p_index, vis) {
            vis[n_index] = true   
            //console.log(adj[n_index])
            for(const it of adj[n_index]) {
                if (vis[it] == false) {
                    if(is_cycle(it, n_index, vis) == true) return true
                } else if(it != p_index) {
                    return true
                }
            }
            return false
        }
        
        
        for(let i = 0; i<V; i++) {
            if(vis[i] == false) {
                if(is_cycle(i, -1, vis) == true) return true
            }
        }
        return false
    }
}
