//User function Template for javascript

/**
 * @param {number[][]} mat
 * @param {number} n
 * @param {number} k
 * @returns {number}
*/


class MinHeap {
    constructor(cap) {
       this.heap_size = 0
       this.harr = []
       this.capacity = cap
    }
    left(i) {
       return 2 * i + 1 
    }
    right(i) {
        return 2 * i + 2
    }
    parent(i) {
        return Math.floor((i - 1)/2)
    }
    
    heapify(i) {
        let l = this.left(i)
        let r = this.right(i)
        let smallest = i
        //console.log(this.harr)
        //console.log([l,r])
        
        if(l < this.heap_size && this.harr[i].val > this.harr[l].val) smallest = l
        if(r < this.heap_size && this.harr[smallest].val > this.harr[r].val) smallest = r
        //console.log([smallest, i, l, r])
        //console.log(this.harr[i].val, this.harr[l].val)
        if(smallest != i) {
            let tmp = this.harr[i]
            this.harr[i] = this.harr[smallest]
            this.harr[smallest] = tmp
            this.heapify(smallest)
        }
        
    }
}

class Node {
    constructor(r,c,val) {
        this.r = r
        this.c = c
        this.val = val
    }
}

class Solution {
  kthSmallest(mat,n,k){
    //code here
    let heap = new MinHeap(n*n)
    for(let i = 0; i<n; i++) {
        heap.harr[i] = new Node(0,i,mat[0][i])
        heap.heap_size++
        
    }
    
    if(k == 1) return heap.harr[0].val
    
    let hr
    for(let i=1; i<=k; i++) {
        hr = heap.harr[0]
        //console.log([hr.r, n-1])
        let next_val = hr.r < n - 1 ? mat[hr.r + 1][hr.c] : Number.MAX_VALUE
        heap.harr[0] = new Node(hr.r + 1, hr.c, next_val)
        heap.heapify(0)
    }
    
    //console.log(heap.harr)
    return hr.val
    
  }
  
}
