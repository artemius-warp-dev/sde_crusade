//User function Template for javascript


/**
 * @param {Node} node
 * @return {number}
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
    /* Should return data of middle node. If linked list is empty, then  -1*/
    getMiddle(node)
    {
        //your code here
        if(node == null)
            return -1
        
        let it_1 = node
        let it_2 = it_1
        let middle
        
        
        while(it_2 != null && it_2.next != null) {
            it_1 = it_1.next
            it_2 = it_2.next.next
            
            
        }
        return it_1.data
    }
}
