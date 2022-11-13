//User function Template for javascript

/**
 * @param {Node} head1
 * @param {Node} head2
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
    //Function to merge two sorted linked list.
    sortedMerge(head1, head2)
    {
        //your code here
        if(head1 == null) return head2
        if(head2 == null) return head1
        
        let l1 = head1
        let l2 = head2
        
        if(l1.data > l2.data) {
            let swap = l1
            l1 = l2
            l2 = swap
        }
        let res = l1
        while(l1 != null && l2 != null) {
            let tmp = null
            while(l1 != null && l1.data <= l2.data) {
                tmp = l1
                l1 = l1.next
            }
            tmp.next = l2
            let swap = l1
            l1 = l2
            l2 = swap
        }
        return res
    }
}
