//User function Template for javascript

/**
 * @param {Node} head
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
    //Function to remove a loop in the linked list.
    removeLoop(head)
    {
        //your code here
        if(head.next == null || head.next.next == null)
            return head
        let slow = head.next
        let fast = slow.next
        while(fast && fast.next) {
            if(slow == fast) break
            slow = slow.next
            fast = fast.next.next
            
        }
        
       //console.log(slow, head)
       if(slow == fast) {
           
           slow = head
           if(slow != fast) {
               while(slow.next != fast.next) {
                slow = slow.next
                fast = fast.next
               }
               fast.next = null
           } else {
              while(slow != fast.next) {
                  fast = fast.next
              } 
              fast.next = null
           }
           
       }
    }
    
}
