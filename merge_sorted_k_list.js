//User function Template for javascript

/**
 * @param {Node[]} arr
 * @param {number} N
 * @returns {Node}
*/



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
            this.harr[0] = this.harr[this.heap_size - 1]
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
        while (i !== 0 && this.harr[this.parent(i)].data > this.harr[i].data) {
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
        if (l < this.heap_size && this.harr[l].data < this.harr[i].data) smallest = l;
        if (r < this.heap_size && this.harr[r].data < this.harr[smallest].data) smallest = r;
        if (smallest != i) {
            let temp = this.harr[i];
            this.harr[i] = this.harr[smallest];
            this.harr[smallest] = temp;
            this.MinHeapify(smallest);
        }
    }
}
/*
class Node{
    constructor(data){
        this.data = data;
        this.next = null;
    }
}
*/

class Solution {
    //Function to merge K sorted linked list.
    mergeKLists(arr, K)
    {
        //your code here
        let heap = new MinHeap(K)
        for(let i = 0; i<K; i++) {
            heap.insertKey(arr[i])
        }
        let dummy = new Node(null)
        let last = dummy
        while(heap.heap_size > 0) {
            let curr = heap.extractMin()
            
            last.next = curr
            last = last.next
            //break
            if(curr.next != null) {
                heap.insertKey(curr.next)
            }
            //console.log(heap.harr[0])
        }
        //console.log(heap.harr)
        return dummy.next
        
    }
}
