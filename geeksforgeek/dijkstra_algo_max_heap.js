
//User function Template for javascript

class Node {
    constructor(s, distance) {
        this.source = s
        this.distance = distance
    }
}

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
        return min
        
    }
    /**
     * @param {number} k
     */
    insertKey(k) {
        // Your code here.
        if(this.heap_size == this.capacity) {
            console.log("LIMIT")
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
        while (i !== 0 && this.harr[this.parent(i)].distance > this.harr[i].distance) {
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
        if (l < this.heap_size && this.harr[l].distance < this.harr[smallest].distance) smallest = l;
        if (r < this.heap_size && this.harr[r].distance < this.harr[smallest].distance) smallest = r;
        if (smallest != i) {
            let temp = this.harr[i];
            this.harr[i] = this.harr[smallest];
            this.harr[smallest] = temp;
            this.MinHeapify(smallest);
        }
    }
}







/**
 * @param {number} V
 * @param {number[][][]} Adj
 * @param {number} S
 * @return {number[]}
 */
class Solution
{
    //Function to find the shortest distance of all the vertices
    //from the source vertex S.
    dijkstra(V,Adj,S)
    {
        //code here
        let distance = new Array(V).fill(Number.MAX_VALUE)
        distance[S] = 0
        let pq = new MinHeap(V*V)
        pq.insertKey(new Node(S, 0))
        
        while(pq.heap_size != 0) {
            //console.log(pq.harr[0])
            let node = pq.extractMin()
            for(const it of Adj[node.source]) {
                if(it[1] + distance[node.source] < distance[it[0]]) {
                    distance[it[0]] = distance[node.source] + it[1]
                    pq.insertKey(new Node(it[0], distance[it[0]]))
                }
            }
        }
        //console.log(pq.harr)
        
        //console.log(distance)
        return distance
        
    }
    
}
