//User function Template for javascript

class Node {
    constructor(r,c, distance) {
        this.row = r
        this.col = c
        this.distance = distance
    }
}

class MinHeap { 
    constructor() {
        this.heap_size = 0;
        // this.capacity = cap;
        this.harr = new Array();
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
        // // Your code here.
        // if(this.heap_size == this.capacity) {
        //     return
        // }
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
        if (l < this.heap_size && this.harr[l].distance < this.harr[i].distance) smallest = l;
        if (r < this.heap_size && this.harr[r].distance < this.harr[smallest].distance) smallest = r;
        if (smallest != i) {
            let temp = this.harr[i];
            this.harr[i] = this.harr[smallest];
            this.harr[smallest] = temp;
            this.MinHeapify(smallest);
        }
    }
}




//User function Template for javascript

/**
 * @param {number[][]} arr
 * @param {number} n
 * @returns {number}
 */

// TLE ERROR. TODO FIX

class Solution 
{
    //Function to return the minimum cost to react at bottom
	//right cell from top left cell.
    minimumCostPath(grid, n)
    {
        //your code here
        let distance = new Array(n).fill(Number.MAX_VALUE).map(() => new Array(n).fill(Number.MAX_VALUE))
        distance[0][0] = grid[0][0]
        let visited = new Array(n).fill(false).map(() => new Array(n).fill(false))
        let pq = new MinHeap()
        pq.insertKey(new Node(0, 0, grid[0][0]))
        
        while(pq.heap_size != 0) {
            //console.log(pq.heap_size)
            
            let node = pq.extractMin()
            let row = node.row
            let col = node.col
            if(row == n-1 && col == n-1) return distance[row][col]
            
            if(visited[row][col]) continue
            
            
            
            visited[row][col] = true
            //let dis = node.distance
                    
            let dr = [-1, 0, 1, 0]
            let dc = [0, 1, 0, -1]

            for(let i =0; i<4; i++) {
                let nrow = row + dr[i]
                let ncol = col + dc[i]
                
                 if(nrow >=0 && ncol >= 0 && nrow < n && ncol < n && !visited[nrow][ncol]) {
                    
                    distance[nrow][ncol] = Math.min(distance[nrow][ncol], distance[row][col] + grid[nrow][ncol])
                    pq.insertKey(new Node(nrow, ncol, distance[nrow][ncol]))
                 } else {
                     continue
                 }
            
            }
        }
        //console.log(pq.harr)
        //console.log(pq.harr, pq.heap_size)
        //console.log(distance)
        return distance[n-1][n-1]
        
        
        
    }
}
