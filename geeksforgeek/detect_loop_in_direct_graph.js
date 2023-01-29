// User function Template for javascript

/**
 * @param {number} V
 * @param {number[][]} adj
 * @returns {boolean}
 */

// BFS - Topological SORT - Kahn's Algo
let isCycleBFS = (V, adj) => {
    // create indegree array
    let indegree = new Array(V);
    for(let i = 0; i < V; i++) {
        indegree[i] = 0;
    }
            
            
    // get all indegrees
    for (let i = 0; i < adj.length; i++) {
        for (let j = 0; j < adj[i].length; j++) {
            indegree[adj[i][j]]++;
        }
    }
    
    // console.log(indegree);
    
    let q = []
    
    // push all nodes with indegree 0 to queue
    for (let i = 0; i < V; i++) {
        if (indegree[i] == 0) q.push(i);
    }
    
    // console.log("q", q);
    
    let ans = [];
    
    let count = 0;
    while(q.length) {
        
        // pop the front element and push it to queue
        let front = q.shift();
        ans.push(front);
        
        count++;
        // get all the neighbours of front
        // reduce their indegree by 1
        // if indegree 0 - add to queue
        for (let j = 0; j < adj[front].length; j++) {
            let neighbour = adj[front][j];
            indegree[neighbour]--;
            if (indegree[neighbour] == 0) {
                q.push(neighbour);
            }
        }
    }
    
    if (count == V) return 0;
    else return 1;
}


//callstack memory error on big adj matrix
class Solution {
    // Function to detect cycle in a directed graph.
    isCyclic(V, adj) {
        // code here
        let vis = new Array(V).fill(false)
        let dfs_vis = new Array(V).fill(false)
        
        let check_cycle = function(node, vis, dfs_vis) {
           
            if(dfs_vis[node]) return true
            if(vis[node]) return false
            
            vis[node] = true
            dfs_vis[node] = true
            
           let children = adj[node];

           for (const it of children) {
                if (check_cycle(it, vis, dfs_vis))
                          return true
           }
            
            dfs_vis[node] = false
            return false
        }
        
        for(let i = 0; i<V; i++) {
            if(check_cycle(i, vis, dfs_vis)) return true
        }
        return false
        
    }
}
