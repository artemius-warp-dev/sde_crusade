//User function Template for javascript

/**
 * @param {Node} head
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
    //Function to reverse a linked list.
    reverseList(head)
    {
        //your code here
        let next = null
        let cur = head
        let prev = null
        while(cur != null) {
            next = cur.next
            cur.next = prev
            prev = cur
            cur = next
            
        }
        head = prev
        return head
    }
    
}
