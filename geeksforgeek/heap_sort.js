/**
 *
 * heapify
 * @param {number[]} arr
 * @param {number} n
 * @param {number} i
 *
 * buildHeap
 * @param {number[]} arr
 * @param {number} n 
 * 
 * heapSort
 * @param {number[]} arr
 * @param {number} n
 */
class Solution
{
    
    parent(i){
      return Math.floor((i - 1) / 2); 
    }
    left(i){ 
      return (2 * i + 1);
    }
    right(i){ 
      return (2 * i + 2); 
    }
    
    
    //Heapify function to maintain heap property.
    heapify(arr,n,i)
    {
    //code here
    let l = this.left(i);
    let r = this.right(i);
    let largest = i;
    if (l < n && arr[l] > arr[largest]) largest = l;
    if (r < n && arr[r] > arr[largest]) largest = r;
    
    if (largest != i) {
        let temp = arr[i];
        arr[i] = arr[largest];
        arr[largest] = temp;
        this.heapify(arr, n, largest);
    }
    
        
    }
    
    //Function to build a Heap from array.
    buildHeap(arr,n) {
      for(let i=Math.floor(n/2-1); i>=0; i--) {
          //console.log(i)
          this.heapify(arr, n, i)
      }
      return arr
    
    }
    
    //Function to sort an array using Heap Sort.
    heapSort(arr,n)
    {
    //code here
    let heap = this.buildHeap(arr, n)
    //console.log(heap)
    for(let i=n-1; i>0; i--) {
        let tmp = heap[0]
        heap[0] = heap[i]
        heap[i] = tmp
        this.heapify(heap, i, 0)
    }
    
    let res = heap
    return res
    
        
    }
    
}
