//User function Template for javascript

/**
 * @param {Node} head
 * @param {number} n
 * @returns {number}
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
    //Function to find the data of nth node from the end of a linked list
    getNthFromLast(head, n)
    {
        //your code here
        let slow = head
        let fast = head
        let cnt = 0
       
            while(cnt < n-1){
                cnt = cnt  + 1
                fast = fast.next
                 if(fast == null)
                    return -1
            }
            
            while(fast.next != null) {
                fast = fast.next
                slow = slow.next
            }
            
            return slow.data
    }
    
}
