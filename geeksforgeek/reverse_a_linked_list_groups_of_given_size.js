 * @param {Node} node
 * @param {number} k
 * @return {Node} node
*/


/*
class Node{
    constructor(data){
        this.data = data;
        this.next = null;
    }
}

let head;
This is method only submission.
You only need to complete the below method.
*/

class Solution {
    /* Should do this in-place without altering the nodes' values.*/
    reverse(node, k)
    {
        //your code here
        if(head == null || head.next == null || k == 1)
            return head
            
            
        let dummy = new Node(null)
        dummy.next = head
        let cur = dummy, next = dummy, prev = dummy
      
        let count = 0
        while(cur.next != null) {
            count++
            cur = cur.next
        }
        
        while(count >= k) {
            cur = prev.next
            next = cur.next
            this.reverse_group(prev, cur, next, k)
   
            
            prev = cur
            count = count - k
            
        }
        
        if(count != 0) {
            cur = prev.next
            next = cur.next
            this.reverse_group(prev, cur, next, count)
        }
  
        
        return dummy.next
        
    }
    
    
    reverse_group(prev, cur, next, k) {
        for(let i = 1; i < k; i++) {
            cur.next = next.next
            next.next = prev.next
            prev.next = next
            next = cur.next
        }        
    }
}
