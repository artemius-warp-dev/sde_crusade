//User function Template for javascript

/**
 * @param {Node} head1
 * @param {Node} head2
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
    //Function to find intersection point in Y shaped Linked Lists.
    intersectPoint(head1, head2)
    {
        //your code here
        
        if (head1 == null || head2 == null)
            return -1
        
        let a = head1
        let b = head2
        while(a != b) {
            a = (a == null) ? head2 : a.next
            b = (b == null) ? head1 : b.next
        }
        
        if(a != null)
          return a.data
        else
          return -1
    }
}
