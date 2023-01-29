//User function Template for javascript

/**
 * @param {number} n
 * @param {number[} a
 * @returns {number}
*/


const CHARS=26

class Graph {
    constructor(v) {
        this.v = v
        this.adj = new Array(v).fill().map(() => new Array())
        this.in = new Array(v).fill(0)
    }
    
    add_edge(v,w) {
        this.adj[v].push(w)
        this.in[w]++
    }
    
    is_eulerian_cycle() {
        if(this.is_sc() == false)
            return false
            
            
        for(let i=0; i<this.v; i++) {
            if(this.adj[i].length != this.in[i])
                return false
        }
        
        return true
    }
    
    dfs(v, visited) {
        visited[v] = true
        for(let it of this.adj[v]) {
            if(!visited[it]) {
                this.dfs(it, visited)
            }
        }
    }
    
    
    get_transponse() {
        let graph = new Graph(this.v)
        for(let v=0; v<this.v; v++) {
            for(let it of this.adj[v]) {
                graph.adj[it].push(v)
                graph.in[v]++
            }
        }
        return graph
    }
    
    is_sc() {
        let visited = new Array(this.v).fill(false)
        let n
        for(n =0; n<this.v; n++) {
            if(this.adj[n].length > 0)
                break
        }
        
        this.dfs(n, visited)
        
        for(let i=0; i<this.v; i++) {
            if(this.adj[i].length > 0 && visited[i] == false)
                return false
        }
        
        let graph = this.get_transponse()
        
        visited.fill(false)
        
        graph.dfs(n, visited)
        
        for(let i=0; i<this.v; i++) {
            if(this.adj[i].length > 0 && visited[i] == false)
                return false
        }        
        
      return true  
    }
    
}


class Solution {
    
    isCircle(n, a)
    {
        //your code here
        let graph = new Graph(CHARS)
        
        for(let i =0; i<n; i++) {
            let s = a[i]
            let v_index = s[0].charCodeAt(0) - 'a'.charCodeAt(0)
            let w_index = s[s.length - 1].charCodeAt(0) - 'a'.charCodeAt(0)
            graph.add_edge(v_index, w_index)
        }
        
       // console.log(graph)
       if(graph.is_eulerian_cycle()) return 1
       else return 0
    }
}
