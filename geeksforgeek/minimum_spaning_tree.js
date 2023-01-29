
//working c++ code

// public:
// 	//Function to find sum of weights of edges of the Minimum Spanning Tree.
//     int spanningTree(int V, vector<vector<int>> adj[])
//     {
//         priority_queue<pair<int, int>, vector<pair<int, int>>, greater<pair<int, int>>> pq;
//         vector<int> vist(V, 0);
//         int sum = 0;
//         // {wt, node}
//         pq.push({0, 0}); // we are only keeping track of weight and node and not the parent
//         // as we only have to retrun the sum and not the edges in MST.
        
//         while(!pq.empty()){
//             auto it = pq.top();
//             pq.pop();
            
//             int wt = it.first;
//             int node = it.second;
            
//             if(vist[node] == 1) continue; // if already visisted
//             vist[node] = 1; // else mark it visited only after popping out of the min-heap
            
//             sum += wt; // add the edge weight to sum
            
//             for(auto it: adj[node]){
//                 int adjNode = it[0];
//                 int edgwt = it[1];
                
//                 if(!vist[adjNode])
//                     pq.push({edgwt, adjNode});
//             }
//         }
        
//         return sum;
//     }

//TLE error js code in GFG
//User function Template for javascript



class MinHeap { 
    constructor(cap) {
        this.heap_size = 0;
        this.capacity = cap;
        this.harr = new Array(cap);
    }
    parent(i){
      return Math.floor((i - 1) / 2); 
    }
    left(i){ 
      return (2 * i + 1);
    }
    right(i){ 
      return (2 * i + 2); 
    }

    // You need to write code for below three functions
  /**
   * @return {number}
   */
    extractMin() {
        // Your code here.
        let min = -1
        if(this.heap_size > 0) {
            min = this.harr[0]
            this.harr[0] = this.harr[this.heap_size-1]
            this.harr.pop()
            this.heap_size--
            this.MinHeapify(0)
        }
        //console.log(min.weight)
        return min
        
    }
    /**
     * @param {number} k
     */
    insertKey(k) {
        // Your code here.
        if(this.heap_size == this.capacity) {
            console.log("cap")
            return
        }
        this.decreaseKey(this.heap_size, k)
        this.heap_size++
        
    }
    
    /**
     * @param {number} i
     */
    deleteKey(i) {
        // Your code here.
        if(this.heap_size <= 0 || this.heap_size <= i) {
            return
        }
        this.decreaseKey(i, 0)
        this.extractMin()
    }

    // Decrease key operation, helps in deleting the element
    decreaseKey(i,  new_val) {
        this.harr[i] = new_val;
        while (i !== 0 && this.harr[this.parent(i)].weight > this.harr[i].weight) {
            let temp = this.harr[i];
            this.harr[i] = this.harr[this.parent(i)];
            this.harr[this.parent(i)] = temp;
            i = this.parent(i);
        }
    }

    /* You may call below MinHeapify function in
      above codes. Please do not delete this code
      if you are not writing your own MinHeapify */
    MinHeapify(i) {
        let l = this.left(i);
        let r = this.right(i);
        let smallest = i;
        if (l < this.heap_size && this.harr[l].weight < this.harr[i].weight) smallest = l;
        if (r < this.heap_size && this.harr[r].weight < this.harr[smallest].weight) smallest = r;
        if (smallest != i) {
            let temp = this.harr[i];
            this.harr[i] = this.harr[smallest];
            this.harr[smallest] = temp;
            this.MinHeapify(smallest);
        }
    }
}

class MinHeapNode {
    constructor(v, w) {
        this.verticle = v
        this.weight = w
    }
}


/**
 * @param {number[][]} arr
 * @param {number} v
 * @param {number} e
 * @returns {number}
*/
class Solution {
    spanningTree(arr, v, e) {
        // code here
        let min_heap = new MinHeap(v*e)
        let vis = new Array(v).fill(-1)
        let adj = new Array(v).fill().map(() => new Array())
        
        //[ [  [1, 5], [2, 1] ], [ [0, 5], [2, 3]  ], [ [1, 3], [0, 1] ] ]
        
        
       // console.log(arr)
       let i = -1
       while(i++ < e-1) {
          let u = arr[i][0]
          let v = arr[i][1]
          let w = arr[i][2]
        //   let t1 = new Array()
        //   let t2 = new Array()
        //   t1.push(v)
        //   t1.push(w)
        //   t2.push(u)
        //   t2.push(w)
          adj[u].push([v,w])
          adj[v].push([u,w])
       }
        
        //console.log(arr[0][1], arr[0][2])
        min_heap.insertKey(new MinHeapNode(0,0))
        let sum = 0
        while(min_heap.heap_size != 0) {
            //console.log(min_heap.harr)
            let node = min_heap.extractMin()
            //console.log(node)
            if(vis[node.verticle] == 1) continue
            vis[node.verticle] = 1
            //console.log(node)
            //console.log(sum)
            sum+=Number(node.weight)
            //console.log(adj[node.verticle])
            for(let it of adj[node.verticle]) {
                //console.log(it, node.verticle)
                if(vis[it[0]] == -1) {
                    //console.log(it, node.verticle)
                    min_heap.insertKey(new MinHeapNode(Number(it[0]), Number(it[1])))
                }
            }
            
            //break
        }
        //console.log(vis)
        return sum
        
        
        
    }
}
