// User function Template for javascript

/**
 * @param {number} V
 * @param {number[][]} adj
 * @returns {number[]}
*/
class Solution {
    // Function to return Breadth First Traversal of given graph.
    bfsOfGraph(V, adj) {
        // code here
        let visited = new Array(V+1).fill(false)
        let queue = []
        let list = []
        visited[0] = true
        queue.push(0)
        //console.log(adj)
        while(queue.length != 0) {
            let m = queue.shift()
            list.push(m)
            //console.log(adj[m])
            for(const i of adj[m]) {
                //console.log(i)
                if(visited[i] == false) {
                    visited[i] = true
                    queue.push(i)
                }
            }
        }
        return list
    }
}
