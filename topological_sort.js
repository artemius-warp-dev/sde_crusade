
//User function Template for javascript

/**
 * @param {number} V
 * @param {number[][]} adj
 * @returns {number[]}
*/
class Solution 
{
    //Function to return list containing vertices in Topological order.
    topoSort(V, adj)
    {
        // code here
        let vis = new Array(V).fill(false)
        let stack = []
        
        
        let find_topo_sort = function(node, vis, adj, st) {
            vis[node] = true
            for(const it of adj[node]) {
                if(!vis[it]) find_topo_sort(it, vis, adj, st)
            }
            st.push(node)
        }
        
        for(let i=0; i<V; i++) {
            if(!vis[i]) find_topo_sort(i, vis, adj, stack)
        }
        
        let order = []
        while(stack.length != 0) {
            order.push(stack.pop())
        }
        
        //console.log(order)
        
        return order
        
    }
}
