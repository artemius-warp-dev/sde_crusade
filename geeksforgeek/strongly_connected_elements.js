
//User function Template for javascript

/**
 * @param {number[][]} arr
 * @param {number} v
 * @param {number} e
 * @returns {number}
*/
class Solution {
    
    
    
    
    
    kosaraju(arr, v, e) {
        // code here
        let vis = new Array(v).fill(false)
        let adj = new Array(v)
        for(let i=0; i<v; i++) {
            adj[i] = new Array()
        }
        
        
        let edje_list = arr
        for(let i =0; i<e; i++) {
            let source = edje_list[i][0]
            let dest = edje_list[i][1]
            adj[source].push(dest)
        }
        //console.log(adj)
        
    
        

        
       
        
        let stack = this.topoSort(adj.length,adj, vis)
        
        //console.log(stack)
        
        let count = 0
        let dfs = function(node, vis, transponse) {
            vis[node] = true
    
            for(const it of transponse[node]) {
                if(!vis[it]) {
                    dfs(it, vis, transponse)
                }
            }
        }
        
        
        //create a temp arraylist to store the reverse 
        let transponse = new Array()
       
        for(let i =0;i<v;i++){
            transponse[i] = new Array()
        }
        
         //console.log(arr)
        for(let i=0; i<v;i++) {
            vis[i] = false
            for(let it of adj[i]) {
                transponse[it].push(i)
            }
        }
        
        // console.log(transponse)
       
        //console.log(stack)
        while(stack.length != 0) {
          let node = stack.pop()
          if(!vis[node]) {
              count++
              dfs(node, vis, transponse)
          }
        }
        
        return count
        
    }
    
    
    
    //Function to return list containing vertices in Topological order.
    topoSort(V, adj, vis)
    {
        // code here
       
        let stack = []
        
        let find_topo_sort = function(node, vis, adj, st) {
            vis[node] = true
            //console.log(adj[node])
            for(const it of adj[node]) {
                if(!vis[it]) find_topo_sort(it, vis, adj, st)
            }
            st.push(node)
        }
        
        for(let i=0; i<V; i++) {
            if(!vis[i]) find_topo_sort(i, vis, adj, stack)
        }
        
        return stack
    }
    
}
