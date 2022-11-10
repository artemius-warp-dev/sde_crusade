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
    //Function to sort a linked list of 0s, 1s and 2s.
    segregate(head)
    {
        //your code here
        let count = [0,0,0]
        let ptr = head
        while(ptr != null) {
            count[ptr.data]++
            ptr = ptr.next
        }
        
        
        let i = 0
        ptr = head
        while(ptr != null) {
            if(count[i] == 0) {
                i++
            } else {
                ptr.data = i
                count[i]--
                ptr = ptr.next
            }
        }
        return head
    }
}
