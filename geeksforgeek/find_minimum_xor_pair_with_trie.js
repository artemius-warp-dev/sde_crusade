//User function Template for javascript

/**
 * @param {number[]} arr
 * @param {number} n
 * @return {number}
 */
 
 class Node {
     constructor() {
        this.links = new Array(2).fill(null)
        this.value = 0
     }
     
     containsKey(bit) {
         return (this.links[bit] != null)
     }
     
     get(bit) {
         return this.links[bit]
     }
     
     put(bit, node) {
         this.links[bit] = node
     }
 }
 
 class Trie {
     constructor() {
         this.root = new Node()
     }
     
     insert(num) {
         let node = this.root
         for(let i=63; i>=0; i--) {
             let bit = (num >> i) & 1
             if(!node.containsKey(bit)) {
                 node.put(bit, new Node())
             }
             node = node.get(bit)
         }
         node.value = num
     }
     
     getMin(num) {
         let node = this.root
         for(let i=63; i>=0; i--) {
             let bit = (num >> i) & 1 
             if(node.containsKey(bit)) {
                 //console.log("same")
                 node = node.get(bit)
             } else if(node.containsKey(1-bit)) {
                 //console.log("oposite")
                 node = node.get(1-bit)
             }
         }
         //console.log(node.value, num)
         return node.value ^ num
     }
     
 }

class Solution {
    minxorpair(arr,n){
      //code here
      let trie = new Trie
      //for(let it of arr) trie.insert(it)
      trie.insert(arr[0])
      //console.log(trie.root.links)
      
      let min  = Number.MAX_VALUE
      
      for(let i =1; i<n; i++) {
          
          min = Math.min(min, trie.getMin(arr[i]))
          trie.insert(arr[i])
          //console.log(min, [t1, t2])
          
          
      }
      return min
      
    }
}

