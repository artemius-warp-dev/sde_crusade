//User function Template for javascript

/*LINKED LIST NODE
class Node {
  constructor(x){
    this.data = x;
    this.next = null;
    this.bottom = null;
  }
}
*/

/**
 * @param {Node} head
 * @return {Node}
*/

class Solution {
    
    flatten(head)
    {
        //your code here
        if(head == null || head.next == null) {
            return head
        }
        
        head.next = this.flatten(head.next)
        head = this.merge(head, head.next)
        return head
    }
    
    merge(l1, l2) {
        let tmp = new Node(null)
        let res = tmp
        while(l1 != null && l2 != null) {
            if(l1.data < l2.data) {
                tmp.bottom = l1
                tmp = tmp.bottom
                l1 = l1.bottom
            } else {
                tmp.bottom = l2
                tmp = tmp.bottom
                l2 = l2.bottom
            }
        }
        if(l1 != null) tmp.bottom = l1
        else tmp.bottom = l2
        
        return res.bottom
    }   
}
