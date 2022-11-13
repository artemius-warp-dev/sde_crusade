//User function Template for javascript

/**
 * @param {Node} head
 * @param {number} k
 * @returns {Node}
*/

/*
class Node{
    constructor(data){
        this.data = data;
        this.next = null;
    }
}
*/

class Solution {
    //Function to rotate a linked list.
    rotate(head, k)
    {
        //your code here
        if(head == null || head.next == null || k == 0)
            return head
        
        let ptr = head
        let len = 1
        while (ptr.next != null) {
            ptr = ptr.next
            len++
        }
        
        ptr.next = head
        
        
        
        
        let i = 1
        while (i <= k) {
            ptr = ptr.next
            i++
        }
        
        //console.log(ptr)
        
        head = ptr.next
        ptr.next = null
        return head
        
        
    }
}
