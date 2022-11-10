//User function Template for javascript

/**
 * @param {Node} head
 * @returns {boolean}
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
    //Function to check if the linked list has a loop.
    detectLoop(head)
    {
        //your code here
        if(head == null || head.next == null)
            return false
        
        
        let slow = head
        let fast = head
        while(fast.next != null && fast.next.next != null) {
            slow = slow.next
            fast = fast.next.next
            if(slow == fast)
                return true
        }
        return false
    }
    
}
