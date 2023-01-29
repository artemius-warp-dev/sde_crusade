// class Solution
// {
//     priority_queue<int, vector<int>, greater<int>> pqmin;
//     priority_queue<int> pqmax;
//     int size=0;
//     public:
//     //Function to insert heap.
//     void insertHeap(int &x)
//     {
//         if(size==0){
//             pqmax.push(x);size++;
//         }
//         else{
//             if(x < pqmax.top()){
//                 pqmax.push(x);size++;
//             }
//             else{
//                 pqmin.push(x);size++;
//             }
//         }
//         balanceHeaps();
//     }
    
//     //Function to balance heaps.
//     void balanceHeaps()
//     {
//         if((int)(pqmin.size()-pqmax.size())>1){
//             int t = pqmin.top();
//             pqmin.pop();
//             pqmax.push(t);
//         }
//         else if((int)(pqmax.size()-pqmin.size())>1){
//             int t = pqmax.top();
//             pqmax.pop();
//             pqmin.push(t);
//         }
//     }
    
//     //Function to return Median.
//     double getMedian()
//     {
//         if(size%2==0){
//             double ans = (pqmin.top()+pqmax.top())/2;
//             return ans;
//         }
//         if(pqmin.size()>pqmax.size()){
//             return pqmin.top(); 
//         }
//         else{
//             return pqmax.top();
//         }
//     }
// };


//User function Template for javascript
// TLE error on GFG

class MaxHeap { 
    constructor(cap) {
        this.heap_size = 0;
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
    extractMax() {
        // Your code here.
        let max = -1
        if(this.heap_size > 0) {
            max = this.harr[0]
            this.harr[0] = this.harr[this.heap_size-1]
            this.harr.pop()
            this.heap_size--
            this.MaxHeapify(0)
        }
        return max
        
    }
    /**
     * @param {number} k
     */
    insertKey(k) {
        // Your code here.
        this.increaseKey(this.heap_size, k)
        this.heap_size++
        
    }
    
    top() {
        return this.harr[0]
    }
    
    /**
     * @param {number} i
     */
    deleteKey(i) {
        // Your code here.
        if(this.heap_size <= 0 || this.heap_size <= i) {
            return
        }
        this.increaseKey(i, 0)
        this.extractMax()
    }

    // Decrease key operation, helps in deleting the element
    increaseKey(i,  new_val) {
        this.harr[i] = new_val;
        while (i !== 0 && this.harr[this.parent(i)] < this.harr[i]) {
            let temp = this.harr[i];
            this.harr[i] = this.harr[this.parent(i)];
            this.harr[this.parent(i)] = temp;
            i = this.parent(i);
        }
    }

    /* You may call below MinHeapify function in
      above codes. Please do not delete this code
      if you are not writing your own MinHeapify */
    MaxHeapify(i) {
        let l = this.left(i);
        let r = this.right(i);
        let largest = i;
        if (l < this.heap_size && this.harr[l] > this.harr[largest]) largest = l;
        if (r < this.heap_size && this.harr[r] > this.harr[largest]) largest = r;
        if (largest != i) {
            let temp = this.harr[i];
            this.harr[i] = this.harr[largest];
            this.harr[largest] = temp;
            this.MaxHeapify(largest);
        }
    }
}





class MinHeap { 
    constructor(cap) {
        this.heap_size = 0;
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
    
    top() {
        return this.harr[0]
    }
    
    /**
     * @param {number} k
     */
    insertKey(k) {
        // Your code here.
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
        while (i !== 0 && this.harr[this.parent(i)] > this.harr[i]) {
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
        if (l < this.heap_size && this.harr[l] < this.harr[smallest]) smallest = l;
        if (r < this.heap_size && this.harr[r] < this.harr[smallest]) smallest = r;
        if (smallest != i) {
            let temp = this.harr[i];
            this.harr[i] = this.harr[smallest];
            this.harr[smallest] = temp;
            this.MinHeapify(smallest);
        }
    }
}




/**
 *
 *insertHeap
 * @param {number} x
 *
 *getMedian
 * @return {float}
 */
class Solution
{
    constructor() {
        this.min_heap = new MinHeap()
        this.max_heap = new MaxHeap()
        this.size = 0
    }
    
    // Function to insert heap
    insertHeap(x)
    {
        // add your code here
        if(this.size == 0) {
            this.max_heap.insertKey(x)
            this.size++
        } else {
            if(x < this.max_heap.top()) {
                this.max_heap.insertKey(x)
                this.size++
            } else {
                this.min_heap.insertKey(x)
                this.size++
            }
        }
        
        this.balanceHeaps()
        
    }
    
     // Function to balance Heaps
    balanceHeaps()
    {
       // add your code here
       let t = null
       if((this.min_heap.heap_size - this.max_heap.heap_size) > 1) {
           t = this.min_heap.extractMin()
           this.max_heap.insertKey(t)
           
       } else if((this.max_heap.heap_size - this.min_heap.heap_size) > 1) {
           t = this.max_heap.extractMax()
           this.min_heap.insertKey(t)
       }
       
    }
    
    // function to getMedian
    getMedian()
    {
        // add your code here
        if(this.size % 2==0 ) {
            let ans = (this.max_heap.top() + this.min_heap.top())/2
            return ans
        }
        if(this.min_heap.heap_size > this.max_heap.heap_size) {
            return this.min_heap.top()
        } else {
            return this.max_heap.top()
        }
        
    }
    
}
