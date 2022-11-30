
//User function Template for javascript

/**
 * @param {string} str
 * @returns {string}
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
        //console.log(this.heap_size, this.capacity)
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
        while (i !== 0 && this.harr[this.parent(i)].freq < this.harr[i].freq) {
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
        if (l < this.heap_size && this.harr[l].freq > this.harr[i].freq) smallest = l;
        if (r < this.heap_size && this.harr[r].freq > this.harr[smallest].freq) smallest = r;
        if (smallest != i) {
            let temp = this.harr[i];
            this.harr[i] = this.harr[smallest];
            this.harr[smallest] = temp;
            this.MinHeapify(smallest);
        }
    }
}

class Node {
    constructor(c,freq) {
        this.char = c
        this.freq = freq
    }
}


class Solution {
    rearrangeString(str){
        //code here
        let n = str.length
        let heap = new MinHeap(n)
        let cache = new Map()
        let char_arr = str.split("")
        for(let i =0; i<n; i++) {
            if(cache.has(char_arr[i])) {
                let count = cache.get(char_arr[i])
                count++
                cache.set(char_arr[i], count )
            } else {
                cache.set(char_arr[i], 1)
            }
        }
        
        for (let i = 97; i <= 122; i++) {
            let c = String.fromCharCode(i)
            if(cache.has(c)) {
                let freq = cache.get(c) 
                heap.insertKey(new Node(c, freq))
            }
        }
        
        //console.log(heap.heap_size)
        let prev = new Node('#', -1)
        let res = ""
        

        
        while(heap.heap_size != 0) {
            let node = heap.extractMin()
            //console.log(heap.heap_size)
            //console.log(node)
            res = res + node.char
            //console.log(res)
            if(prev.freq > 0) {
                heap.insertKey(prev)
            }
            //console.log(node)
            node.freq--
            //console.log(node)
            prev = node
        }
        //console.log(heap.heap_size)
        //console.log(heap.harr)
        //console.log(res)
        if(res.length != n) return 0
        else return 1
        
    }
}
