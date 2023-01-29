//User function Template for javascript

/**
 * @param {string[]} dict
 * @param {number} N
 * @param {number} K
 * @return {string}
 */
class Solution
{
  findOrder(dict,N,K){
      //code here
      let adj = new Array(K)
      for(let i=0; i<K; i++) {
          adj[i] = new Array()
      }
      //console.log(adj)
      for(let i=0; i<N-1; i++) {
         
          let s1 = dict[i]
          let s2 = dict[i+1]
          let len = Math.min(s1.length, s2.length)
          for(let j = 0; j<len; j++) {
              if(s1[j] != s2[j]) {
                  //console.log(s1[j].charCodeAt(0) - 'a'.charCodeAt(0), s2[j].charCodeAt(0) - 'a'.charCodeAt(0))
                  //console.log(s1[j].charCodeAt(0) - 'a'.charCodeAt(0))
                  //console.log(adj[s1[j].charCodeAt(0) - 'a'.charCodeAt(0)])
                  adj[s1[j].charCodeAt(0) - 'a'.charCodeAt(0)].push(s2[j].charCodeAt(0) - 'a'.charCodeAt(0))
                  break
              }
          }
      }
      //console.log(adj)
      let stack = this.topoSort(K, adj)
      let ans = ""
        while(stack.length != 0) {
            let ch = stack.pop()
            ans = ans + String.fromCharCode(ch + 'a'.charCodeAt(0))
        }
        
       // console.log(ans)
      return ans
  }
  
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
        
       
        
        return stack
        
    }
    
}
